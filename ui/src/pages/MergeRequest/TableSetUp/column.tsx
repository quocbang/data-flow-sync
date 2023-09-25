import { ColumnsType } from "antd/es/table";
import { Tag } from "antd";
import { mergeRequest, state } from "../type";

const setTag = (s: state) => {  
  let color = '';
  switch (s) {
    case 1: {
      color = 'green'
      break;
    }
    case 2: {
      color = 'gold'
      break;
    }
    case 3: {
      color = 'red'
      break;
    }    
    case 4: {
      color = 'yellow'
      break;
    }    
    default: {
      color = ''
      break;
    }
  }    
  return (
    <>
      <Tag color={color} key={s}>{state[s]}</Tag>
    </>
  )
}

export const subColumn: ColumnsType<mergeRequest> = [    
  {
    title: 'Title',
    dataIndex: 'title',
    key: 'title',
  },
  {
    title: 'State',
    key: 'state',
    dataIndex: 'state',
    render: (_, { state }) => (
      <>        
        {state && setTag(state) }                  
      </>
    ),
  },
  {
    title: 'File Affected',
    key: 'fileAffected',
    dataIndex: 'fileAffected'
  },
  {
    title: 'Author',
    key: 'author',
    dataIndex: 'author'
  },
  {
    title: 'orderDate',
    dataIndex: 'orderDate',
    key: 'orderDate'
  },
];