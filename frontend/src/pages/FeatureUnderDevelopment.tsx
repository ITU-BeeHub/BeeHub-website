import React from 'react';
import BeakerIcon from '../components/icons/BeakerIcon'; // You can replace with any icon relevant to the feature

const FeatureUnderDevelopment: React.FC<{ featureName: string }> = ({ featureName }) => {
  return (
    <div className="min-h-screen flex flex-col bg-[#F5FDFD] h-screen overflow-y-auto">
      {/* Scrollable container */}
      <div className="flex-1 overflow-y-scroll px-4 py-8">
        <div className="max-w-4xl w-full mx-auto p-8 bg-white rounded-lg shadow-lg">
          <div className="mb-8">
            <BeakerIcon className="h-20 w-20 text-[#FDC003] mx-auto" />
          </div>
          <h1 className="text-4xl font-bold text-[#0372CE] text-center mb-6">
            {featureName} Under Development
          </h1>
          <p className="text-lg text-gray-600 text-center mb-10">
            We are working hard to bring {featureName} to life. Stay tuned for future updates and exciting features!
          </p>
          <p className="text-gray-600 text-center">
            Meanwhile, feel free to explore other sections of BeeHub.
          </p>
        </div>
      </div>
    </div>
  );
};

export default FeatureUnderDevelopment;
