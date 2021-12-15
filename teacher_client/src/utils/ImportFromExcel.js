import XLSX from 'xlsx'

export function extract_from_excel(file) {
  return new Promise((resolve, reject) => {
      let reader = new FileReader()
      reader.onerror = () => {
        reject()
      }
      reader.onload = () => {
        let f = XLSX.read(reader.result)
        let sheets = f.SheetNames.map(name => f.Sheets[name])
        resolve(sheets.map(sheet => XLSX.utils.sheet_to_json(sheet)).flat(1))
      }
      reader.readAsArrayBuffer(file)
    }
  )
}
