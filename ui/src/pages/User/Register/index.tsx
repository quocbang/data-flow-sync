import React from 'react';
import { useEmotionCss } from '@ant-design/use-emotion-css';
import { Helmet, history } from '@umijs/max';
import { useIntl, SelectLang,FormattedMessage, useModel } from '@umijs/max';
import Form, { useForm } from 'antd/es/form/Form';
import FormItem from 'antd/es/form/FormItem';
import { Button, Input } from 'antd';
import Password from 'antd/es/input/Password';
import { registerAccount } from '@/services/api/api';
import { decodeToken } from '..';
import { flushSync } from 'react-dom';
import { setToken } from '@/utils/Cookies';
import { MailOutlined, UserOutlined, LockOutlined  } from '@ant-design/icons';
import { checkIsEmailExisted, checkIsUserExisted } from '@/services/graphql/api/api';

const Lang = () => {
  const langClassName = useEmotionCss(({ token }) => {
    return {
      width: 50,
      height: 50,
      lineHeight: '42px',
      position: 'fixed',
      right: 16,
      borderRadius: token.borderRadius,
      ':hover': {
        backgroundColor: token.colorBgTextHover,
      },
    };
  });

  return (
    <div className={langClassName} data-lang>
      {SelectLang && <SelectLang />}
    </div>
  );
};

const Register: React.FC = () => {
  const { setInitialState } = useModel('@@initialState');
  const [ form ] = useForm();
  const intl = useIntl()
  const containerClassName = useEmotionCss(() => {
      return {
        display: 'flex',
        flexDirection: 'column',
        height: '100vh',
        overflow: 'auto',
        backgroundImage:
          "url('https://mdn.alipayobjects.com/yuyan_qk0oxh/afts/img/V-_oS6r-i7wAAAAAAAAAAAAAFl94AQBr')",
        backgroundSize: '100% 100%',
      };
  });   
  const handleRegister = async (value: API.RegisterParams) => {      
    const data = await registerAccount({ ...value })
    if (data.token) {
      const userInfo = decodeToken(data.token)
      if (userInfo) {
        flushSync(() => {                
          setInitialState(() => ({          
            token: data.token,
            currentUser: userInfo,
          }));
        });
      }
      setToken(data.token, userInfo.exp || 0)
      const urlParams = new URL(window.location.href).searchParams;
      history.push(urlParams.get('redirect')  || '/')
    }
  }
  const isNicknameExisted = async (userID: string): Promise<boolean> => {
    if (userID == "") {
      return false
    }
    const resp = await checkIsUserExisted(userID)    
    if (resp.account?.length !== 0) {
      return true
    }
    return false
  }
  const isEmailExisted = async (email: string): Promise<boolean> => {
    if (email === "") {
      return false
    }
    const resp = await checkIsEmailExisted(email)
    if (resp.account?.length !== 0) {
      return true
    }else {
      return false
    }
  }
  return (
      <div className={containerClassName}>
          <Helmet>
            <title>
              {intl.formatMessage({
                id: 'menu.register',
                defaultMessage: 'Register',
              })}                                
            </title>
          </Helmet>
          <Lang></Lang>
          <div
            style={{
              display: 'flex',
              flexDirection: 'column',                
              padding: '100px 0',      
              alignItems: 'center',                
              textAlign: 'center'                          
            }}  
          >
            <Form                              
              form={form}
              name='register'              
              onFinish={ async (value) => {
                await handleRegister(value as API.RegisterParams)
              }}
              style={{ minWidth: 350 ,  maxWidth: '75vw' }}
              scrollToFirstError                                
            >
              <h1>{intl.formatMessage({
                id: 'menu.register',
                defaultMessage: 'Register',
              })}</h1>
              <FormItem
                name='email'
                validateTrigger={[
                  "onBlur"                  
                ]}
                rules={[
                  {
                    required: true,
                    message: 'Please input your E-mail!',
                  }, 
                  {
                    type: 'email',
                    message: 'The input is not valid E-mail!',                                        
                  },
                  {                                             
                    validator(rule, value, callback) {                                                                                   
                      isEmailExisted(value).then((result) => {                        
                        if (result === true) {
                          callback("email already existed")
                        }else {
                          callback()
                        }
                      })
                    },                    
                  },                                                  
                ]}                                        
                hasFeedback                                
              >
                <Input 
                  prefix={<MailOutlined></MailOutlined>}
                  placeholder={intl.formatMessage({
                    id: 'pages.register.email',
                    defaultMessage: 'E-mail',
                  })}                  
                >                      
                </Input>
              </FormItem>
              <FormItem
                name='name'                
                // TODO: use graphql to check nick name
                validateTrigger={"onBlur"}
                rules={[       
                  {
                    required: true,
                    message: 'Please input your nick name!'
                  },             
                  {
                    validator(_, value, callback) {                      
                      isNicknameExisted(value).then((result) =>{
                        if (result === true) {
                          callback("Nickname already existed")
                        }else {
                          callback()
                        }
                      })                      
                    },
                  }           
                ]} 
                hasFeedback                                 
              >
                <Input
                  prefix={<UserOutlined />}                  
                  placeholder={intl.formatMessage({
                    id: 'pages.register.nickName',
                    defaultMessage: 'Nickname',
                  })}
                >                  
                </Input>
              </FormItem>
              <FormItem
                name='password'                
                // TODO: use graphql to check nick name
                rules={[ 
                  {
                    required: true,
                    message: 'Please input your password!',                      
                  },        
                  {
                    min: 8,
                    message: 'Your Password is too weak please choose a password longer than 8 characters'
                  }
                ]} 
                hasFeedback                 
              >
                <Password
                  prefix={<LockOutlined />}
                  placeholder={intl.formatMessage({
                    id: 'pages.register.password',
                    defaultMessage: 'Password',
                  })}
                >                  
                </Password>
              </FormItem>
              <FormItem
                name='confirmPassword'                
                dependencies={['password']}
                hasFeedback
                rules={[
                  {
                    required: true,
                    message: 'Please confirm your password!',
                  },
                  ({ getFieldValue }) => ({
                    validator(_, value) {
                      if (!value || getFieldValue('password') === value) {
                        return Promise.resolve();
                      }else {
                        return Promise.reject(new Error('The new password that you entered do not match!'))
                      }
                    }
                  })
                ]}                  
              >
                <Password
                  prefix={<LockOutlined />}
                  placeholder={intl.formatMessage({
                    id: 'pages.register.confirmPassword',
                    defaultMessage: 'Confirm Password',
                  })}
                >                  
                </Password>
              </FormItem>
              <FormItem>
                <Button style={{ width: '100%' }} type='primary' htmlType='submit'>
                  {intl.formatMessage({
                    id: 'menu.register',
                    defaultMessage: 'Register',
                  })}
                </Button>
              </FormItem>
              <div
                style={{
                  marginBottom: 24,
                }}
              >
                <a
                  style={{
                    float: 'right',
                  }}
                  href='/user/login'
                >
                  <FormattedMessage id="pages.login.loginMessage" defaultMessage='Already have an account? Sign In'></FormattedMessage>
                </a>
              </div>
            </Form>
          </div>
      </div>    
  );
}

export default Register;