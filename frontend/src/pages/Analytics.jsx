import React from 'react'
import InvoiceAnalytics from '../components/analytics/InvoiceAnalytics'
import InvoiceQuarterly from '../components/analytics/InvoiceQuarterly'
import AdvertisementDuration from '../components/analytics/AdvertisementDuration'
import TotalSalary from '../components/analytics/TotalSalary'
import NumEmployees from '../components/analytics/NumEmployees'

function Analytics() {
  return (
    <div className='pt-20'>
      <h1 className="text-4xl font-bold mb-8 mt-10 underline text-center text-gray-800">Business Analytics Dashboard</h1>
    <InvoiceAnalytics />
    <InvoiceQuarterly />
    <AdvertisementDuration />
    <TotalSalary />
    <NumEmployees />
  </div>
  )
}

export default Analytics