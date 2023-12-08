import React from 'react';
import { BarChart, Bar, XAxis, YAxis, Tooltip, Legend, ResponsiveContainer } from 'recharts';

const QuarterlyInvoiceChart = ({ data }) => {
  const formatQuarter = (quarter) => `Q${quarter}`;
  const formatTooltip = (value, name, props) => [`${value}`, `${formatQuarter(props.payload.quarter)}`];

  return (
    <ResponsiveContainer width="60%" height={400}>
      <BarChart data={data} margin={{ top: 20, right: 30, left: 20, bottom: 5 }}>
        <XAxis dataKey="quarter" tickFormatter={formatQuarter} label={{ value: 'Quarter', position: 'insideBottom', offset: -5 }} />
        <YAxis dataKey="totalAmount" label={{ value: 'Total Amount', angle: -90, position: 'insideLeft', offset: -5 }} />
        <Tooltip formatter={formatTooltip} />

        <Bar dataKey="totalAmount" fill="blue" name="Total Amount" />
      </BarChart>
    </ResponsiveContainer>
  );
};

export default QuarterlyInvoiceChart;
