import React from 'react';
import {
  BarChart,
  Bar,
  XAxis,
  YAxis,
  Tooltip,
  Legend,
  ResponsiveContainer,
  CartesianGrid,
} from 'recharts';

const InvoiceChart = ({ data }) => {
  return (
    <ResponsiveContainer width="50%" height={400}>
      <BarChart data={data} margin={{ top: 20, right: 30, left: 20, bottom: 5 }}>
        <CartesianGrid strokeDasharray="3 3" />
        <XAxis dataKey="paymentStatus" label={{ value: 'Payment Status', position: 'insideBottom', offset: -5 }} />
        <YAxis dataKey="amount" label={{ value: 'Amount', angle: -90, position: 'insideLeft', offset: -5 }} />
        <Tooltip />

        {/* Bar for 'Paid' */}
        <Bar dataKey="amount" fill="green" />

      </BarChart>
    </ResponsiveContainer>
  );
};

export default InvoiceChart;
