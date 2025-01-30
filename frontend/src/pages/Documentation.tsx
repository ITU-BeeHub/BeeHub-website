import React from 'react';
import { Link } from 'react-router-dom'; // For internal linking

import PickaxeIcon from '../components/icons/PickaxeIcon';
import ArchiveIcon from '../components/icons/ArchiveIcon';

const Documentation: React.FC = () => {
  return (
    <div className="container mx-auto px-4 py-8">
      <h1 className="text-4xl font-bold text-center mb-8 text-[#0372CE]">BeeHub Documentation</h1>

      <p className="text-lg text-gray-700 text-center mb-12">
        Learn how to use BeeHub and its various features to manage your university experience efficiently.
      </p>

      {/* Installation Section */}
      <div className="bg-white p-6 rounded-lg shadow-md mb-12">
        <h2 className="text-2xl font-bold text-[#0372CE] text-center mb-4">Installation</h2>
        <p className="text-gray-600 mb-4 text-center">
          Follow the steps below to install BeeHub on your device.
        </p>

        <div className="text-left space-y-6">
          {/* Windows */}
          <div>
            <h3 className="text-xl font-semibold text-gray-800 mb-2">For Windows</h3>
            <p className="text-gray-600">
              1. Download the BeeHub installer for Windows from <a href="/" className="text-[#0372CE] hover:underline">here</a>.
            </p>
            <p className="text-gray-600">
              2. Run the installer and follow the on-screen instructions to complete the setup.
            </p>
            <p className="text-gray-600">
              3. Once installed, launch BeeHub from the Start menu or desktop shortcut.
            </p>

            {/* Windows SmartScreen Warning */}
            <div className="mt-4">
              <img src="/images/smartscreen.png" alt="Windows SmartScreen" className="mx-auto" />
              <p className="text-gray-600 text-center mt-2">
                If you encounter a Windows SmartScreen warning, click on <strong>More info</strong>, and then select <strong>Run anyway</strong>.
              </p>
            </div>
          </div>

          {/* Disclaimer Section */}
          <div className="mt-8 p-4 bg-yellow-100 border-l-4 border-yellow-500 text-yellow-700">
            <p className="font-semibold">Disclaimer:</p>
            <p className="text-gray-700">
              You may need to trust this installer/application as Windows might give a warning during installation.
              This is because we could not sign our application with an EV (Extended Validation) certificate. Unfortunately,
              we don't have the budget to afford this certificate at the moment. Rest assured, the application is safe to use.
            </p>
          </div>

          {/* macOS */}
          <div>
            <h3 className="text-xl font-semibold text-gray-800 mb-2">For macOS</h3>
            <p className="text-gray-600">
              1. Download the BeeHub installer for macOS from <a href="/" className="text-[#0372CE] hover:underline">here</a>.
            </p>
            <p className="text-gray-600">
              2. Open the downloaded `.dmg` file and drag BeeHub to the Applications folder.
            </p>
            <p className="text-gray-600">
              3. Launch BeeHub from the Applications folder or by using Spotlight search.
            </p>

            {/* macOS Terminal Command */}
            <div className="mt-4">
              <img src="/images/terminal.jpg" alt="macOS Terminal" className="mx-auto" />
              <p className="text-gray-600 text-center mt-2">
                If macOS prevents the app from running, open Terminal and run the following command:
              </p>
              <pre className="bg-gray-200 p-2 rounded-md text-sm">xattr -cr /Applications/BeeHub.app</pre>
            </div>
          </div>

          {/*
          //{/* Linux //}
          <div>
            <h3 className="text-xl font-semibold text-gray-800 mb-2">For Linux</h3>
            <p className="text-gray-600">
              1. Download the BeeHub package for Linux from <a href="/" className="text-[#0372CE] hover:underline">here</a>.
            </p>
            <p className="text-gray-600">
              2. Install the package using your preferred package manager (e.g., `dpkg`, `rpm`).
            </p>
            <p className="text-gray-600">
              3. Launch BeeHub from your application launcher.
            </p>
          </div>
          */}
        </div>
      </div>

      {/* Features Section */}
      <div className="grid grid-cols-1 lg:grid-cols-2 gap-12">
        {/* BeePicker Section */}
        <div className="bg-white p-6 rounded-lg shadow-md">
          <div className="flex justify-center mb-6">
            <PickaxeIcon className="h-16 w-16 text-[#FDC003]" />
          </div>
          <h2 className="text-2xl font-bold text-[#0372CE] text-center mb-4">BeePicker</h2>
          <p className="text-gray-600 mb-4 text-center">
            BeePicker is your personal assistant for selecting courses. Filter courses, avoid conflicts, and pick your schedule.
          </p>
          <Link to="/beepicker" className="text-[#0372CE] font-semibold hover:underline text-center block">
            Learn more about BeePicker
          </Link>
        </div>

        {/* BeeArchive Section */}
        <div className="bg-white p-6 rounded-lg shadow-md">
          <div className="flex justify-center mb-6">
            <ArchiveIcon className="h-16 w-16 text-[#FDC003]" />
          </div>
          <h2 className="text-2xl font-bold text-[#0372CE] text-center mb-4">BeeArchive</h2>
          <p className="text-gray-600 mb-4 text-center">
            Access and manage your past courses, exams, and notes with BeeArchive. Easily browse and download archived files.
          </p>
          <Link to="/beearchive" className="text-[#0372CE] font-semibold hover:underline text-center block">
            Learn more about BeeArchive
          </Link>
        </div>
      </div>

      <div className="pb-8"></div> {/* Single padding at bottom instead of multiple breaks */}
    </div>
  );
};

export default Documentation;