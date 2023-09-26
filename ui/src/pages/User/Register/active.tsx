import Alert from "antd/es/alert/Alert"
import React, { useState } from 'react'
import Modal from "antd/es/modal/Modal"
import { GetCode, verifyAccount } from "@/services/api/api";
import { decodeToken } from "..";
import { useModel, useIntl } from "@umijs/max";
import { flushSync } from "react-dom";
import { setToken } from "@/utils/cookies";
import { ProFormCaptcha } from "@ant-design/pro-components";

export declare type ActiveOptions = {
    isUnspecifiedUser?: boolean;
}

const CheckAndActiveUser: React.FC<ActiveOptions> = (data) => {             
  const intl = useIntl();
  const [open, setOpen] = useState(false)           
  const handleCancel = () => {
    setOpen(false)
  }

  // set code to state.
  const [ code, setCode ] = useState("")
  const handleSetCode = (value: string) => {      
    setCode(value)      
  }
  
  // get code in state and send api to verify.
  const [loading, setLoading] = useState(false); 
  const { setInitialState } = useModel('@@initialState')
  const handleOk  = async () => {
      setLoading(true)                
      var otp: API.VerifyAccountParams = {
        otp: code,
      } 
      try {
        const data = await verifyAccount(otp)
        if (data.token) {
          const userInfo = decodeToken(data.token)
          if (userInfo) {
            flushSync(() => {
              setInitialState(() => ({
                token: data.token,
                currentUser: userInfo,
              }));
            })
          }
          // replace token on Cookie.
          setToken(data.token, userInfo.exp || 0)
        }
      } catch {
        setLoading(false)
      }      
  }
  const handleOpenToggleClick = () => {
      setOpen(true)
  }    
  const handleGetCode = async () => {
    await GetCode()
  }    
  const handleClose = (): boolean => {
    return true;
  }

  // notify when user is unspecified user
  if (data.isUnspecifiedUser && !open) {
    return (
      <>
        <Alert
          onClick={handleOpenToggleClick}
          style={{
              cursor: 'pointer',
            }}
            banner
            message="Your account hasn't been activated yet! Click here to activate."
        />        
      </>
    );
  }        

  return (
      <>
        <Modal
          title="Active Account"
          open={open || false}
          onOk={handleOk}
          onCancel={handleCancel}
          destroyOnClose={handleClose()}                      
          confirmLoading={loading}
        >           
          <ProFormCaptcha
            onGetCaptcha={handleGetCode}
            captchaTextRender={(timing, count) => {
              if (timing) {
                return `${count} ${intl.formatMessage({
                  id: 'pages.getCaptchaSecondText',
                  defaultMessage: '获取验证码',
                })}`;
              }
              return intl.formatMessage({
                id: 'pages.login.phoneLogin.getVerificationCode',
                defaultMessage: '获取验证码',
              });
            }}
            name="code"              
            onChange={(e: any) => handleSetCode(e.target.value)}            
            rules={[
              {
                required: true,
                message: 'Please enter your code'
              },
              {
                len: 6,
                message: 'Verify code are 6 characters'
              }                
            ]}                        
          >
          </ProFormCaptcha>
        </Modal>           
      </>
  );
}

export default CheckAndActiveUser;