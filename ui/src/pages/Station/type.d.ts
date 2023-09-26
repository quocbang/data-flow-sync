declare namespace STATION {
  type Station = {
    ID?: string
    subCompany?: number
    factory?: string
    departmentID?: string
    alias?: string
    serialNumber?: number
    description?: string
    devices?: number[]
  }  

  type StationList = {
    ID?: string
    size?: string    
    createAt?: string
    modified?: string
    owner?: string
  }
}