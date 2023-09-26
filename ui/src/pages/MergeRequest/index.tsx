import React, { useState } from "react";
import { FloatButton, Modal, Table } from "antd";
import { mergeRequest } from "./type";
import { subColumn } from "./TableSetUp/column";
import { ColumnsType } from "antd/es/table";
import { Link } from "@umijs/max";
const mergeRequestPath = "/merge-request"

const mergeRequests: mergeRequest[] = [
  {id: 'MergeRequest1', title: 'add new limitary hour file', state: 1, fileAffected: "limitary-hour.yaml", author: 'quocbang',  orderDate: '2023-08-26'},
  {id: 'MergeRequest2', title: 'add new station id in station group', state: 2, fileAffected: "KV-P8300-PLY-CUT-1.yaml", author: 'quocbang',  orderDate: '2023-08-26'},
  {id: 'MergeRequest3', title: 'delete station', fileAffected: "KV-P8322-PLY-CUT-1.yaml", state: 3, author: 'quocbang',  orderDate: '2023-08-26'},
  {id: 'MergeRequest3', title: 'delete station', fileAffected: "KV-P8322-PLY-CUT-1.yaml", state: 3, author: 'quocbang',  orderDate: '2023-08-26'},  
];

const MergeRequest: React.FC = () => {
  const [ open, setOpen ] =useState(false)
  const handleOpenMergeRequest = () => {      
    handleListMergeRequest()
    setOpen(true)
  }
  
  const [ selectMergeRequestRows, setMergeRequestRows ] = useState<mergeRequest[]>([])
  const handleListMergeRequest = () => {    
    setMergeRequestRows(mergeRequests)    
  }  

  // show MR button only
  if (!open) {      
      return (
        <FloatButton.Group
          shape="circle"          
        >                
          <FloatButton onClick={handleOpenMergeRequest} badge={{ count: mergeRequests.length, overflowCount: 999 }} />
          <FloatButton.BackTop visibilityHeight={0} />
        </FloatButton.Group> 
      );
  }

  const setUpColumns: ColumnsType<mergeRequest> = [
    {
      title: 'ID',
      dataIndex: 'id',
      key: 'id',
      render: (mergeRequestID) => {                          
        return (
          <Link to={`${mergeRequestPath}/${mergeRequestID}/detail`} onClick={() => setOpen(false)}>{mergeRequestID}</Link>
        )
      }
    },
    ...subColumn
  ]

  return (
    <Modal
      title="Merge Request Notify"
      centered
      open={open}
      onOk={() => setOpen(false)}
      okButtonProps={{ hidden: true }}
      onCancel={() => setOpen(false)}
      width={1500}
    >
      {selectMergeRequestRows.length > 0 && (
        <Table 
          columns={setUpColumns}
          dataSource={selectMergeRequestRows}        
        >
        </Table>
      )}
    </Modal>
  );
}

export default MergeRequest