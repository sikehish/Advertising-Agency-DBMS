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
  
  
  export function AdsDurationTable({ads}) {
    return (
      <Table className="bg-gray-100 border-gray-600 text-black">
        <TableCaption>Quarterly Invoice Stats</TableCaption>
        <TableHeader>
          <TableRow>
            <TableHead>Duration</TableHead>
            <TableHead>No. of Ads</TableHead>
            <TableHead>Total Cost</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          {ads.map((ad, ind) => (
            <TableRow key={ind}>
              <TableCell>{ad.durationCategory}</TableCell>
              <TableCell>{ad.numAds}</TableCell>
              <TableCell>{ad.totalCost}</TableCell>
            </TableRow>
          ))}
        </TableBody>
        {/* <TableFooter>
          <TableRow>
            <TableCell colSpan={3}>Total</TableCell>
            <TableCell className="text-right">{ads[0].amount+ads[1].amount}</TableCell>
          </TableRow>
        </TableFooter> */}
      </Table>
    )
  }
  