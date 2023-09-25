export enum state {
  'ADD' = 1,
  'MODIFY',
  'DELETE',
  'RENAME'
}

export interface mergeRequest {
    id?: string,
    title?: string,
    state?: state,    
    fileAffected?: string,
    author?: string,
    orderDate?: string
}

export function getKeyByValue(key: number): string | undefined {
  if (state[key]) {
    return state[key];
  }
  return undefined;
}