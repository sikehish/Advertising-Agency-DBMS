import React, { useEffect, useState } from 'react';
import InvoiceChart from '../charts/InvoiceChart'

const InvoiceAnalytics = () => {
  const [invoiceData, setInvoiceData] = useState([]);

  useEffect(() => {
    fetch('api/v1/analytics/invoices')
      .then(response => response.json())
      .then(data => setInvoiceData(data))
      .catch(error => console.error('Error fetching data:', error));
  }, []);

  console.log(invoiceData)
  return (
    <div>
      <h1>Invoice Data Visualization</h1>
      {invoiceData.length > 0 && <InvoiceChart data={invoiceData} />}
      <section id="invoice-chart"></section>
    </div>
  );
};

export default InvoiceAnalytics;