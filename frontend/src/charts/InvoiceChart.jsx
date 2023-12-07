import React from 'react';
import { Bar } from 'react-chartjs-2';

const InvoiceChart = ({ data }) => {
  const labels = data.map(item => item.payment_status);
  const amounts = data.map(item => item.amount);

  // Chart data
  const chartData = {
    labels: labels,
    datasets: [
      {
        label: 'Amount',
        backgroundColor: 'rgba(75,192,192,0.4)',
        borderColor: 'rgba(75,192,192,1)',
        borderWidth: 1,
        hoverBackgroundColor: 'rgba(75,192,192,0.6)',
        hoverBorderColor: 'rgba(75,192,192,1)',
        data: amounts,
      },
    ],
  };

  // Chart options
  const chartOptions = {
    scales: {
      x: {
        beginAtZero: true,
      },
      y: {
        beginAtZero: true,
      },
    },
  };

  return <Bar data={chartData} options={chartOptions} />;
};

export default InvoiceChart;
