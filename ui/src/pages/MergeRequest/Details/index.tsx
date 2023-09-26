import React, { ReactNode, useEffect, useState } from "react";
import { history, useIntl, useParams } from "@umijs/max";
import { PageContainer } from "@ant-design/pro-components";
import { Button, Card, Col, Row, Tabs, Alert, Tooltip, Timeline, Tag, Tree, TreeProps, Modal, message} from "antd";
import { DownOutlined, ExclamationCircleOutlined } from "@ant-design/icons";
import { useModel } from "@umijs/max";
import { ROLE } from "@/utils/roles";
import { BaseButtonProps } from "antd/es/button/button";
import { parseTreeData } from "./changed";
import { GetStationMergeRequest } from "@/services/api/api";

enum ETabs {
  "conversation" = 1,
  "checks",
  "fileChange"
}

const MergeRequestDetail: React.FC = () => {    
  const init = useIntl()
  const { initialState } = useModel('@@initialState');
  const isNotLeader = initialState?.currentUser?.role !== ROLE.LEADER
  const [ openCloseConfirm, setOpenCloseConfirm ] = useState(false)

  const [ approved, setApproved ] = useState(false)
  const handleShowApproveNotification = () => {
    return (
      <>
        <Alert
          banner
          message={init.formatMessage({
            id: 'pages.mergeRequest.detail.approveRequiredNotification',
            defaultMessage: 'Approve required'            
          })}
        />   
      </>
    );
  }

  const handleDisableMerge = () => {
    // return true if is not leader
    if (isNotLeader) {
      return true
    }
    // return true if have'nt approved
    if (!approved) {
      return true
    }    
    return false
  }

  const [ approveText, setApproveText ] = useState(init.formatMessage({
    id: 'pages.mergeRequest.detail.approve',
    defaultMessage: 'Approve'    
  }))
  const [ approvedType, setApprovedType ] = useState<BaseButtonProps["type"]>("default")
  const handleApproved = () => {    
    if (approved) {
      setApproved(false)
      setApproveText(init.formatMessage({
        id: 'pages.mergeRequest.detail.approve',
        defaultMessage: 'Approve'    
      }))
      setApprovedType("default")
      // TODO: call api to set Approve === false      
    }else {
      setApproved(true)
      setApproveText(init.formatMessage({
        id: 'pages.mergeRequest.detail.approved',
        defaultMessage: 'Approved'    
      }))
      setApprovedType("primary")
      // TODO: call api to set Approve === true
    }
  }

  const handleShowApproveButton = () => {
    const approveButton = <Button disabled={isNotLeader} type={approvedType} onClick={handleApproved}>{approveText}</Button>     
    if (isNotLeader) {          
      return (
        <Tooltip 
          title={init.formatMessage({
            id: 'pages.mergeRequest.detail.approveLeaderRequired',
              defaultMessage: 'Only leader can be approve'    
            })}
        >            
          {approveButton}
        </Tooltip>    
      )
    }
    return approveButton
  }

  const handleShowMergeButton = () => {
    const mergeButton = <Button disabled={handleDisableMerge()} type="primary">{init.formatMessage({
      id: 'pages.mergeRequest.detail.merge',
      defaultMessage: 'Merge'})}</Button>

    // return with tooltip is not leader    
    if (isNotLeader) {
      return (
        <Tooltip 
          title={init.formatMessage({
            id: 'pages.mergeRequest.detail.mergeLeaderRequired',
            defaultMessage: 'Only leader can be merge'
          })}
        >        
          {mergeButton}
        </Tooltip>
      )
    }

    // return with tooltip haven't approved    
    if (!approved) {
      return (
        <Tooltip 
          title={init.formatMessage({
            id: 'pages.mergeRequest.detail.mergeRequireApproveNotification',
            defaultMessage: 'Merge need to approve'
            })}
        >        
          {mergeButton}
        </Tooltip>
      )
    } 

    // return ok
    return mergeButton
  }

  const handleOK = () => {
    // TODO: call API to request close merge request
    message.success("close merge request success")
    history.push("/")
    setOpenCloseConfirm(false)
  }
  const handleCancel = () => {
    setOpenCloseConfirm(false)
  }
  const handleCloseMergeRequestClick = () => {   
    setOpenCloseConfirm(true)
    // confirm
  }

  const [ fileChanged, setFileChanged ] = useState([])
  const handleGetMergeRequestData = async () => {
    const getMErgeRequestResponse = await GetStationMergeRequest(Number(mergeRequestID))
    console.log(getMErgeRequestResponse)
    setApproved(getMErgeRequestResponse.mergeRequestStatus.isApproved)
    setHistoryChanged(getMErgeRequestResponse.mergeRequestInfo.historyChanged)
    setFileChanged(getMErgeRequestResponse.data)
    
  }

  const [ historyChanged, setHistoryChanged ] = useState([])
  const handleParseHistoryChanged = () => {
    let historyChangedData: any[] = []
    historyChanged.forEach((value: string, index: number) => {
      let color = index == historyChanged.length - 1 ? "green" : "gray" // green if is last index
      historyChangedData = [ ...historyChangedData, { color: color, children: value } ]
    })
    return historyChangedData
  }

  const handleOverview = () => {        
    return (
      <>
        <PageContainer
          content={""}
          tags={<Tag color="green" key="mergeRequestStatusTag">Open</Tag>}
          title={
            init.formatMessage({
              id: "pages.mergeRequest.detail.mergeRequest",
              defaultMessage: 'Merge Request'
            }) + " #" + mergeRequestID
          }
        >
          <Row gutter={24}>
            <Col span={24}>
              <Card title={init.formatMessage({
                  id: 'pages.mergeRequest.detail.approve',
                  defaultMessage: 'Approve'    
                })}>
                {!approved && handleShowApproveNotification()}
                {handleShowApproveButton()}                
              </Card>
            </Col>
          </Row>
          <Row gutter={24}>
            <Col span={24}>
              <Card 
                title={
                  init.formatMessage({
                    id: 'pages.mergeRequest.detail.historyChanged',
                    defaultMessage: 'History Changed'
                  })
                }
              >      
                <div>
                  <Timeline                                       
                    items={handleParseHistoryChanged()}
                  />                 
                </div>                         
              </Card>
            </Col>
          </Row> 
          <Row gutter={24}>
            <Col span={24}>
              <Card title={
                init.formatMessage({
                  id: 'pages.mergeRequest.detail.merge',
                  defaultMessage: 'Merge'
                })
              }>              
              <div style={{ display: "flex", justifyContent: "right" }}>
                <Button
                  style={{ marginRight: 10 }}
                  danger
                  onClick={handleCloseMergeRequestClick}
                >{init.formatMessage({
                    id: 'pages.mergeRequest.detail.close',
                    defaultMessage: 'Close',
                  })}
                </Button>
                {handleShowMergeButton()}
              </div>
              </Card>
            </Col>
          </Row>          
        </PageContainer>
        <Modal
          title="Modal"
          open={openCloseConfirm}
          onOk={handleOK}      
          onCancel={handleCancel}          
          okText="Yes"
          
        >
          <p>Are you sure you want to close the merge request?</p>        
        </Modal>
      </>
    );
  }

  const handleShowFileChanged = () => {
    const onSelect: TreeProps['onSelect'] = (selectedKeys, info) => {
      console.log('selected', selectedKeys, info);
    };
    let result: any = []
    fileChanged.forEach((value: any, index: number) => {      
      const file = parseTreeData(value.file);
      const added = parseTreeData(value.added);
      const deleted = parseTreeData(value.deleted);      
      result = [ ...result,  <div style={{ marginBottom: 20, borderTop: '2px solid gray' }} key={index}>
        <Row gutter={24} >
          <Col span={14}>
            <Card
              title={value.file.ID}
              >
              <Tree
                showLine
                switcherIcon={<DownOutlined />}
                defaultExpandedKeys={['0-0-0']}
                onSelect={onSelect}
                treeData={file}                  
              />
            </Card>          
          </Col>
          <Col span={10}>
            <Row>
              <Col span={24}>          
                <Card 
                  title={init.formatMessage({
                    id: 'pages.mergeRequest.detail.added',
                      defaultMessage: 'Added'    
                    })}
                  headStyle={{ backgroundColor: 'lightgreen' }}
                >      
                  <Tree
                    showLine
                    switcherIcon={<DownOutlined />}
                    defaultExpandedKeys={['0-0-0']}
                    onSelect={onSelect}
                    treeData={added}
                  />                        
                </Card>     
              </Col>
            </Row>
            <Row>
              <Col span={24}>
                <Card
                  title={init.formatMessage({
                    id: 'pages.mergeRequest.detail.deleted',
                      defaultMessage: 'Deleted'    
                    })}
                    headStyle={{ backgroundColor: '#ff4d4f' }}
                >
                <Tree
                    showLine
                    switcherIcon={<DownOutlined />}
                    defaultExpandedKeys={['0-0-0']}
                    onSelect={onSelect}
                    treeData={deleted}
                  />
                </Card>   
              </Col>
            </Row>
          </Col>
        </Row>   
      </div>]      
    })
    return (result)
  }

  const handleFileChange = () => {   
    return (
      <>
        {fileChanged !== null && handleShowFileChanged()}
      </> 
    );
  }

  const setupItems = () => {
    return ([
      {
        label: `
          ${init.formatMessage({
            id: 'pages.mergeRequest.detail.tab.conversation',
            defaultMessage: 'Conversation'
          })}
        `,
        key: ETabs.conversation.toString(),
        children: handleOverview(),                
      },
      {
        label: `
          ${init.formatMessage({
            id: 'pages.mergeRequest.detail.tab.check',
            defaultMessage: 'Checks'
          })}
        `,
        key: ETabs.checks.toString(),        
      },
      {
        label: `
          ${init.formatMessage({
            id: 'pages.mergeRequest.detail.tab.fileChanged',
            defaultMessage: 'File Changed'
          })}
        `,
        key: ETabs.fileChange.toString(),   
        children: handleFileChange()     
      }
    ]);
  }

  const { mergeRequestID } = useParams<{ mergeRequestID: string }>();
  // call API to get MR info
  useEffect(() => {
    handleGetMergeRequestData()
  }, [])
  return (
    <>
      <PageContainer
        content={""}
        title={
          init.formatMessage({
            id: 'pages.mergeRequest.detail.title',
            defaultMessage: 'Merge Request Detail'
          })
        }
      >                     
        <Tabs          
          type="card"
          items={setupItems()}                 
        />     
      </PageContainer>
    </> 
  )
}

export default MergeRequestDetail