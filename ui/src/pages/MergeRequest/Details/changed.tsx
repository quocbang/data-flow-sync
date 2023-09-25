import { DataNode } from "antd/es/tree";
import { forEach } from "lodash";

// is parse file, added, deleted response by server to UI tree.
export const parseTreeData = (file: any):DataNode[] => {
  return (parseData(file, 0));
}

// parseData uses recursion to parse data to tree
function parseData(file: any, n: number):DataNode[] {
  if (Array.isArray(file) || typeof file !== 'object') {    
    forEach(file, function(value, key){      
      return [
        {
          title: <strong>{key+":"}</strong>,
          key: key+"-"+n,          
          children: [
            {
              title: value,
              key: value+"-"+n,         
            }
          ]                    
        }
      ]
    })
  }
  var result: any = []
  forEach(file, function(value, key){
    var i = 0
    if (Array.isArray(value) || typeof value === 'object') {      
      result = [...result, {
        title: <strong>{key+":"}</strong>,
        key: key+"-"+n+"-"+i,
        children: parseData(value, n+1)
      }]
    }else {
      result = [...result, {
        title: <strong>{key+":"}</strong>,
        key: key+"-"+n+"-"+i,
        children: [
          {
            title: value,
            key: value+"-"+n+"-"+i,
          }
        ]
      }]
    }
    n++
    i++
  })
  
  return (result)
}