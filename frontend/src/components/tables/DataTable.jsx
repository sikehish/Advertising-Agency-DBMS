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

export function DataTable({ data }) {
  if (!data || data.length === 0) {
    return <p>No data available</p>;
  }

  // Assuming data is an array of objects with key and value properties
  const keys = Object.keys(data[0]);

  return (
    <Table className="bg-gray-100 border-gray-600 text-black">
      <TableCaption>Invoices stats</TableCaption>
      <TableHeader>
        <TableRow>
          {keys.map((key, index) => (
            <TableHead key={index}>{key}</TableHead>
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
