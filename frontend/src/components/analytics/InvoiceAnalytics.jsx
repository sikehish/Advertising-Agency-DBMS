import React, { useEffect, useState } from 'react';
import InvoiceChart from '../charts/InvoiceChart'
import { DataTable } from '../tables/DataTable';

const InvoiceAnalytics = () => {
  const [invoiceData, setInvoiceData] = useState([]);

  useEffect(() => {
    fetch('api/v1/analytics/invoice-payments')
      .then(response => response.json())
      .then(data => setInvoiceData(data))
      .catch(error => console.error('Error fetching data:', error));
  }, []);

  console.log(invoiceData)
  return (<>
      {/* <h1>Invoice Data Visualization</h1> */}
    <div className='flex m-10 justify-center items-center'>
      {invoiceData.length > 0 && <InvoiceChart data={invoiceData} />}
      {invoiceData.length > 0 && <DataTable data={invoiceData} tableName={"Invoice Stats Table"}/>}
    </div>
  </>
  );
};

export default InvoiceAnalytics;