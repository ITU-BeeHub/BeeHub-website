import React from 'react';
import { Link } from 'react-router-dom';
import BeakerIcon from '../components/icons/BeakerIcon'; // Use the BeakerIcon or replace with a relevant icon for 404

const PageNotFound: React.FC = () => {
  return (
    <div className="min-h-screen flex flex-col bg-[#F5FDFD] items-center justify-center text-center px-4">
      <div className="max-w-lg w-full bg-white p-8 rounded-lg shadow-lg">
        <div className="mb-6">
          <BeakerIcon className="h-20 w-20 text-[#FDC003] mx-auto" />
        </div>
        <h1 className="text-4xl font-bold text-[#0372CE] mb-4">404 - Page Not Found</h1>
        <p className="text-lg text-gray-600 mb-6">
          Oops! It looks like the page you're looking for doesn't exist. It might have been moved or removed.
        </p>

        <Link to="/">
          <button className="bg-[#FDC003] text-[#0372CE] font-bold py-3 px-6 rounded-lg shadow-md hover:bg-[#fdc003d9] transition duration-300">
            Go Back Home
          </button>
        </Link>
      </div>
    </div>
  );
};

export default PageNotFound;
