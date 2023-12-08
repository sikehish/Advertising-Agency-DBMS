import React, { useEffect, useState } from 'react';
import InvoiceAnalytics from "./pages/InvoiceAnalytics"
import InvoiceQuarterly from './pages/InvoiceQuarterly';
import AdvertisementDuration from "./pages/AdvertisementDuration"

const App = () => {
 return(
  <div>
  <InvoiceAnalytics />
  <InvoiceQuarterly />
  <AdvertisementDuration />
</div>
 )
};

export default App;
