import { DeleteOutlined, FieldBinaryOutlined, UploadOutlined } from "@ant-design/icons";
import { PageContainer, ProTable } from "@ant-design/pro-components";
import { Link, history } from "@umijs/max";
import { Card, Space, Input, Table, Upload, Button, UploadProps, Modal } from "antd";
import React, { useState } from "react";
import yaml from 'js-yaml';
import { ColumnsType } from "antd/es/table";
import { CreateStationMergeRequest } from "@/services/api/api";
const detailPath = "/station/detail"

const Station: React.FC = () => {  
  const { TextArea } = Input;

  const columns: ColumnsType<STATION.StationList> = [
    {
      title: 'ID',
      dataIndex: 'ID',
      key: 'id',
      render: (id) => {
        return (
          <Link to={detailPath+"/"+id}>{id}</Link>
        );
      },
    },
    {
      title: 'Size',
      dataIndex: 'size',      
      key: 'size',
    },
    {
      title: 'Create Time',
      dataIndex: 'createAt',    
      key: 'createTime',
    },
    {
      title: 'Modified Time',
      dataIndex: 'modified',
      key: 'modified',
    },
    {
      title: 'Owner',
      dataIndex: 'owner',
      key: 'owner'      
    },
  ]
  const data: STATION.StationList[] = [
    {
      ID: "KV-P8322-PLY-CUT-1",
      createAt: "sjdgkdjgdg",
      modified: "andlshslkdhlsdhs",
      size: "10 M",
      owner: "quocbang"   
    }
  ] 
  const rowSelection = {
    onChange: (selectedRowKeys: React.Key[], selectedRows: STATION.StationList[]) => {
      console.log(`selectedRowKeys: ${selectedRowKeys}`, 'selectedRows: ', selectedRows);
    },
    getCheckboxProps: (record: STATION.StationList) => ({
      disabled: record.ID === 'Disabled User', // Column configuration not to be checked
      name: record.ID,
    }),
  };    

  const [ open, setOpen ] = useState(false)
  const handleCancel = () => {    
    setUploadData([])
    setOpen(false)    
    setAliveFileList({}) // reset alive file list
  }
  const handleOK = async () => {        
    let stationRequest: API.CreateStationMergeRequest = {files: []}
    uploadData.forEach((value) => {
      try {
        const jsonObject = yaml.load(value);        
        stationRequest.files = [ ...stationRequest.files, JSON.stringify(jsonObject) ]
      } catch (error) {
        console.error('Error parsing YAML:', error);
      }
    })    
    try {
      const response = await CreateStationMergeRequest(stationRequest)      
      setOpen(false) // close modal
      setAliveFileList({}) // reset alive file list
      setParseError(false)      
      history.push(`/merge-request/${response.mergeRequestID}/detail`)
    }catch (e) {
      console.error("failed to create merge request, error: ", e)
    }
  }

  let allYamlParsed: any[] = []
  const [ uploadData, setUploadData ] = useState<any[]>([])  
  const handleReadData =  async (fileList: any) => {    
    for (let i = 0; i < fileList.length; i++) {
      const yamlData = await readFileAsText(fileList[i].originFileObj)                    
        allYamlParsed = [...allYamlParsed, yamlData]              
    }
    setOpen(true)
    setUploadData(allYamlParsed)
  }

  // alive file is list of file uploaded in each event user click upload button
  // ex:
  //  - First click upload with files A,B,C => aliveFileList = A,B,C.
  //  - Second, click upload With files AND aliveFileList = A,D.
  const [ aliveFileList, setAliveFileList ] = useState<{ [key: string]: boolean }>({})
  const handleUploadChange = async ({ fileList }: any) => {        
    if (fileList.length > 0) {           
      try {                     
        let countDone = 0   
        fileList.map( (file: any) => {
          // get data file with status equal done only.
          if (file.status !== "done") {            
            return
          }else {
            countDone++                        
            if (countDone === fileList.length) {              
              // skip all file not in alive file list.                            
              fileList.map((file: any, index: any) => {                
                if (aliveFileList[file.uid] !== true) {                                    
                  fileList[index].uid = null // set null uids are not exist in alive list
                }                
              })   
              // read data uploaded                         
              handleReadData(fileList.filter((file: any) => file.uid !== null))              
            }
          }          
        })         
      }catch (e) {
        console.error('Error reading YAML:', e);
      }                
    }    
  };

  const readFileAsText = (file: File) => {
    return new Promise<string>((resolve, reject) => {
      const reader = new FileReader();
      reader.onload = () => resolve(reader.result as string);
      reader.onerror = reject;
      reader.readAsText(file);
    });
  };
    
  const handleChangeTextValue = (index: any, newValue: any) => {                
    const updatedData = [...uploadData];
    updatedData[index] = newValue;
    setUploadData(updatedData);
    setParseError(false) // reset error state
  };

  const handleDeleteStationByIndex = (index: any) => {    
    setParseError(false) // reset error state   
    const deleteData = [...uploadData]
    deleteData.splice(index, 1)
    setUploadData(deleteData)    
  }

  const props: UploadProps = {    
    onChange: handleUploadChange,           
    accept: ".yaml",    
    multiple: true,
    beforeUpload(_, fileList) {      
      fileList.map((file: any ) => {
        setAliveFileList( (aliveFileList) => ({...aliveFileList, [file.uid]: true}));        
      })
    },
    defaultFileList: [],
    showUploadList: false,      
  }  
  
  let yamlDataParsed: any = ''
  const [ parseError, setParseError ] = useState(false)
  const handleParseDataToModal = (isError: boolean): string => {    
    if (isError === true) {      
      return yamlDataParsed
    }
    return JSON.stringify(yamlDataParsed, null, 2)
  }
  return (
    <>       
      <Modal
        open={open}
        title="Create Merge Request"
        onCancel={handleCancel}       
        onOk={handleOK}
        okButtonProps={{
          disabled: parseError,          
        }}        
        closable={false}          
        width={1500}           
      >
        <>
          {uploadData.map((value, index) => {
            let isError = false;
            try {              
              yamlDataParsed = yaml.load(value)              
            }catch(e) {              
              isError = true
              yamlDataParsed = e
              if (!parseError) {
                setParseError(true)              
              }
            }            
            return (
              <Card 
                title="Station Parser"
                key={index}                            
                actions={[
                  <Button onClick={() => handleDeleteStationByIndex(index)}>
                    <DeleteOutlined 
                      key={index}                                                      
                    ></DeleteOutlined>
                  </Button>
                ]}
                style={{ backgroundColor: "#f0f0f0", marginBottom: 10 }}
              > 
                <div>
                  <TextArea
                    key={"data-upload-"+index}                    
                    rows={12}                            
                    value={value}                                                
                    style={{ backgroundColor: '#fff', color: '#333', width: '50%' }}
                    onChange={(e) => handleChangeTextValue(index, e.target.value)}
                    typeof="yaml"             
                  />
                  <TextArea                    
                    key={"data-parsed-"+index}
                    rows={12}
                    style={{ backgroundColor: '#272822', color: '#fff', width: '50%' }}
                    value={handleParseDataToModal(isError)}                    
                  />                
                </div>
              </Card>      
            )
          })}
        </>
      </Modal>         
      <PageContainer>
        <Space
          align="center"
          style={{ marginBottom: 16, display: "flex", justifyContent: "right" }}
        >
          <Upload
            {...props}                            
          >      
            <Button
              icon={<UploadOutlined></UploadOutlined>}
              style={{ backgroundColor: 'greenyellow' }}
            >
              Upload
            </Button>    
          </Upload>
        </Space>        
        <Table 
          columns={columns}
          dataSource={data}
          rowSelection={{          
            ...rowSelection,
          }}        
        >
        </Table>
      </PageContainer>
    </>
  );
};

export default Station;