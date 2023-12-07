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
  
  
  export function InvoiceTable({invoices}) {
    return (
      <Table className="bg-gray-100 border-gray-600 text-black">
        <TableCaption>Invoices stats</TableCaption>
        <TableHeader>
          <TableRow>
            <TableHead className="w-[100px]">Invoice count</TableHead>
            <TableHead>Status</TableHead>
            <TableHead className="text-right">Amount</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          {invoices.map((invoice, ind) => (
            <TableRow key={ind}>
              <TableCell className="font-medium">{invoice.invoicesCount}</TableCell>
              <TableCell>{invoice.paymentStatus}</TableCell>
              <TableCell className="text-right">{invoice.amount}</TableCell>
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
  