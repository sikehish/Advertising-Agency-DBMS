import React, { useEffect, useState } from 'react';
import InvoiceChart from '../charts/InvoiceChart'
import { InvoiceTable } from '../components/tables/InvoiceTable';

const InvoiceAnalytics = () => {
  const [invoiceData, setInvoiceData] = useState([]);

  useEffect(() => {
    fetch('api/v1/analytics/invoices')
      .then(response => response.json())
      .then(data => setInvoiceData(data))
      .catch(error => console.error('Error fetching data:', error));
  }, []);

  console.log(invoiceData)
  return (<>
      <h1>Invoice Data Visualization</h1>
    <div className='flex m-10 justify-center items-center'>
      {invoiceData.length > 0 && <InvoiceChart data={invoiceData} />}
      {invoiceData.length > 0 && <InvoiceTable invoices={invoiceData} />}
    </div>
  </>
  );
};

export default InvoiceAnalytics;