import {
  Table,
  TableBody,
  TableCaption,
  TableCell,
  TableFooter,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";

export function DataTable({ data, tableName }) {

  function convertCamelCase(ele){
    ele=ele.replace(/([A-Z])/g, ' $1').replace(/^./, function(str){ return str.toUpperCase(); })
    if(ele.startsWith("Num")) ele=ele.replace("Num ","Number Of ")
    return ele
  }

  if (!data || data.length === 0) {
    return <p>No data available</p>;
  }

  // Assuming data is an array of objects with key and value properties
  const keys = Object.keys(data[0]);

  return (
    <Table className="bg-gray-100 border-gray-600 text-black">
      <TableCaption>{tableName}</TableCaption>
      <TableHeader>
        <TableRow>
          {keys.map((key, index) => (
            <TableHead key={index}>{convertCamelCase(key)}</TableHead>
          ))}
        </TableRow>
      </TableHeader>
      <TableBody>
        {data.map((row, rowIndex) => (
          <TableRow key={rowIndex}>
            {keys.map((key, colIndex) => (
              <TableCell key={colIndex}>{row[key]}</TableCell>
            ))}
          </TableRow>
        ))}
      </TableBody>
    </Table>
  );
}
