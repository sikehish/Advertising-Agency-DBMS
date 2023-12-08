import React from 'react';
import { RadialBarChart, RadialBar, Tooltip, ResponsiveContainer } from 'recharts';

const CustomTooltip = ({ active, payload }) => {
  if (active && payload && payload.length) {
    const { department, totalSalaryExpense } = payload[0].payload;
    return (
      <div style={{ background: 'rgba(255, 255, 255, 0.8)', padding: '10px', border: '1px solid #ccc' }}>
        {`${department} Salary Expense: ${totalSalaryExpense} Rs`}
      </div>
    );
  }

  return null;
};

const SalaryExpenseRadialBarChart = ({ data }) => {
  return (
    <ResponsiveContainer width="100%" height={400}>
      <RadialBarChart cx="50%" cy="50%" innerRadius="10%" outerRadius="80%" data={data}>
        <RadialBar dataKey="totalSalaryExpense" background clockWise={true} />
        <Tooltip content={<CustomTooltip />} />
      </RadialBarChart>
    </ResponsiveContainer>
  );
};

export default SalaryExpenseRadialBarChart;
