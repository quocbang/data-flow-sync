import type { RequestOptions } from '@@/plugin-request/request';
import type { RequestConfig } from '@umijs/max';
import { message, notification } from 'antd';
import { getToken } from './utils/cookies';
import { history } from '@umijs/max';

// Error handling scheme: error type
enum ErrorShowType {
  SILENT = 0,
  WARN_MESSAGE = 1,
  ERROR_MESSAGE = 2,
  NOTIFICATION = 3,
  REDIRECT = 9,
}
// The response data format agreed with the backend
interface ResponseStructure {
  success: boolean;
  data: any;
  errorCode?: number;
  errorMessage?: string;
  showType?: ErrorShowType;
}

// local error
interface ResponseBackendStructure {
  code: number;
  details: string;
}

// swagger error
interface SwaggerErrorStructure {
  code: number;
  message: string;
}


/**
 * @name 错误处理
 * The error handling that comes with pro, you can make your own changes here
 * @doc https://umijs.org/docs/max/request#配置
 */
export const errorConfig: RequestConfig = {  
  // Error handling: umi@3's error handling scheme.
  errorConfig: {
    // error thrown
    errorThrower: (res) => {
      const { success, data, errorCode, errorMessage, showType } =
        res as unknown as ResponseStructure;
      if (!success) {
        const error: any = new Error(errorMessage);
        error.name = 'BizError';
        error.info = { errorCode, errorMessage, showType, data };
        throw error; // 抛出自制的错误
      }
    },
    // Error reception and handling
    errorHandler: (error: any, opts: any) => {
      if (opts?.skipErrorHandler) throw error;
      //The error thrown by our errorThrower
      if (error.name === 'BizError') {
        const errorInfo: ResponseStructure | undefined = error.info;
        if (errorInfo) {
          const { errorMessage, errorCode } = errorInfo;
          switch (errorInfo.showType) {
            case ErrorShowType.SILENT:
              // do nothing
              break;
            case ErrorShowType.WARN_MESSAGE:
              message.warning(errorMessage);
              break;
            case ErrorShowType.ERROR_MESSAGE:
              console.log(error)
              alert(errorInfo)
              message.error(errorMessage);
              break;
            case ErrorShowType.NOTIFICATION:
              notification.open({
                description: errorMessage,
                message: errorCode,
              });
              break;
            case ErrorShowType.REDIRECT:
              console.log(error)
              alert(errorInfo)
              // TODO: redirect
              break;
            default:
              message.error(errorMessage);
          }
        }
      } else if (error.response) {
        if (error.response.status === 401 || error.response.status === 403) {
          notification.warning({
            message: error.response.status,
            description: "session unauthorized, please login again!",
            duration: 3,  
          })
          history.push("/user/login")
          return;
        }
        if (error.response.data) {
          // local error          
          const localError: ResponseBackendStructure | undefined = error.response.data;
          if (localError?.code || localError?.details) {            
            notification.error({
              message: localError.code,
              description: localError.details,              
              duration: 3,              
            })
            return;            
          }
          // swagger error
          const swaggerError: SwaggerErrorStructure | undefined = error.response.data;          
          if (swaggerError?.code && swaggerError?.message) {
            notification.error({
              message: swaggerError.code,
              description: swaggerError.message,
              duration: 3
            })
            return;
          }
        }
        // Axios errors
        // The request was made successfully and the server responded with a status code, but the status code was outside the range of 2xx        
        message.error(`Response status:${error.response.status}`);
      } else if (error.request) {
        // The request was made successfully, but no response was received
        // \`error.request\` is an instance of XMLHttpRequest in the browser,
        // while in node.js is an instance of http.ClientRequest
        message.error('None response! Please retry.');
      } else {
        // Something went wrong while sending the request
        message.error('Request error, please retry.');
      }
    },
  },

  // request interceptor
  requestInterceptors: [
    (config: RequestOptions) => {      
      // Intercept request configuration for personalized processing.
      const url = config?.url;
      const graphql = url?.match("/rest/")
      if (process.env.NODE_ENV === 'production') {          
        // config.baseURL = process.env.REACT_APP_BASE_URL
        // TODO: should get in env
        config.baseURL = 'http://localhost:8888'
      }
      if (graphql) {        
        // config.baseURL = process.env.REACT_APP_GRAPHQL_URL
        // TODO: should get in env
        config.baseURL = 'http://localhost:9999'
      }
      config.headers = {
        'Content-Type': 'application/json',
        'x-data-flow-sync-auth-key': getToken() || '',
        'x-hasura-admin-secret': 'quocbang',
      };
      return { ...config, url };
    },
  ],

  // response interceptor
  responseInterceptors: [
    (response) => {
      // Intercept response data for personalized processing
      const { data } = response as unknown as ResponseStructure;

      if (data?.success === false) {
        message.error('请求失败！');
      }
      return response;
    },
  ],
};
