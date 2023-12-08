import React, { useEffect, useState } from 'react';
import InvoiceChart from '../components/charts/InvoiceChart'
import QuarterlyInvoiceChart from '../components/charts/QuarterlyInvoiceChart';
import { QuarterlyInvoiceTable } from '../components/tables/QuarterlyInvoiceTable';
import { DataTable } from '../components/tables/DataTable';

const InvoiceQuarterly = () => {
  const [invoiceData, setInvoiceData] = useState([]);

  useEffect(() => {
    fetch('api/v1/analytics/invoice-quarterly')
      .then(response => response.json())
      .then(data => setInvoiceData(data))
      .catch(error => console.error('Error fetching data:', error));
  }, []);

  return (<>
      <h1>Quarterly Invoice Data Visualization</h1>
    <div className='flex m-12 justify-center items-center'>
      {invoiceData.length > 0 && <QuarterlyInvoiceChart data={invoiceData} />}
      {invoiceData.length > 0 && <DataTable data={invoiceData} />}
    </div>
  </>
  );
};

export default InvoiceQuarterly;