import {
    Table,
    TableBody,
    TableCaption,
    TableCell,
    TableFooter,
    TableHead,
    TableHeader,
    TableRow,
  } from "@/components/ui/table"
  
  
  export function QuarterlyInvoiceTable({invoices}) {
    return (
      <Table className="bg-gray-100 border-gray-600 text-black">
        <TableCaption>Quarterly Invoice Stats</TableCaption>
        <TableHeader>
          <TableRow>
            <TableHead>Quarter</TableHead>
            <TableHead>Invoice Count</TableHead>
            <TableHead>Total Amount</TableHead>
            <TableHead>Average Amount</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          {invoices.map((invoice, ind) => (
            <TableRow key={ind}>
              <TableCell>{invoice.quarter}</TableCell>
              <TableCell>{invoice.numInvoices}</TableCell>
              <TableCell>{invoice.totalAmount}</TableCell>
              <TableCell>{invoice.averageAmount}</TableCell>
            </TableRow>
          ))}
        </TableBody>
        {/* <TableFooter>
          <TableRow>
            <TableCell colSpan={3}>Total</TableCell>
            <TableCell className="text-right">{invoices[0].amount+invoices[1].amount}</TableCell>
          </TableRow>
        </TableFooter> */}
      </Table>
    )
  }
  