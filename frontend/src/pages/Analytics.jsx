import React from 'react'
import InvoiceAnalytics from '../components/analytics/InvoiceAnalytics'
import InvoiceQuarterly from '../components/analytics/InvoiceQuarterly'
import AdvertisementDuration from '../components/analytics/AdvertisementDuration'
import TotalSalary from '../components/analytics/TotalSalary'
import NumEmployees from '../components/analytics/NumEmployees'

function Analytics() {
  return (
    <div className='pt-20'>
    <InvoiceAnalytics />
    <InvoiceQuarterly />
    <AdvertisementDuration />
    <TotalSalary />
    <NumEmployees />
  </div>
  )
}

export default Analytics