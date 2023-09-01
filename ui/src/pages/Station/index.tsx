import { UploadOutlined } from "@ant-design/icons";
import { PageContainer } from "@ant-design/pro-components";
import { Link } from "@umijs/max";
import { Space, Table, Upload, Button, UploadProps } from "antd";
import React, { useState } from "react";
import yaml from 'js-yaml';

interface DataType {  
  name: string;  
}

const Station: React.FC = () => {
  const initializeColumns = [
    {
      title: "File Name",
      dataIndex: "name",
      icon: "fileTextOutlined",
      render: (id: string) => <Link to={`/station/detail/${id}`}></Link>
    },    
  ]
  const data = [
    {
      key: "1",
      name: "Station 1",
    },
    {
      key: "2",
      name: "Station 2",
    },
    {
      key: "3",
      name: "Station 3",
    },
  ]    
  const rowSelection = {
    onChange: (selectedRowKeys: React.Key[], selectedRows: DataType[]) => {
      console.log(`selectedRowKeys: ${selectedRowKeys}`, 'selectedRows: ', selectedRows);
    },
    getCheckboxProps: (record: DataType) => ({
      disabled: record.name === 'Disabled User', // Column configuration not to be checked
      name: record.name,
    }),
  };  

  const [uploadedData, setUploadedData] = useState<DataType[]>([]);

  const handleUploadChange = async ({ file, fileList }: any) => {
    console.log(fileList)
    if (file.status === 'done') {
      try {
        const yamlData = await readFileAsText(file.originFileObj);
        console.log(yamlData)
        const parsedData = yaml.load(yamlData);
        console.log(parsedData)
        // setUploadedData([...uploadedData, parsedData]);        
      } catch (error) {
        console.error('Error reading or parsing YAML:', error);
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

  const props: UploadProps = {
    onChange: handleUploadChange,
    accept: ".yaml"
  }
  return (
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
          >
            Upload
          </Button>    
        </Upload>
      </Space>
      {/* <Divider></Divider> */}
      <Table
        rowSelection={{          
          ...rowSelection,
        }}
        columns={initializeColumns}
        dataSource={data}
      >
      </Table>
    </PageContainer>
  );
};

export default Station;