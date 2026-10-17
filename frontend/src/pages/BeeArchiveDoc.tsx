import React from 'react';

const BeeArchiveDoc: React.FC = () => {
  return (
    <div className="min-h-screen bg-white py-16 px-4 sm:px-6 lg:px-8">
      <div className="max-w-3xl mx-auto">
        <h1 className="text-4xl font-bold text-[#212121] mb-8">BeeArchive</h1>
        <div className="prose prose-lg">
          <p className="text-gray-600 mb-6">
            BeeArchive is your collaborative platform for accessing and sharing past exam questions.
            Connect with fellow students and build a comprehensive archive of exam materials together.
          </p>

          <h2 className="text-2xl font-semibold text-[#0372CE] mb-4">Key Features</h2>
          <ul className="space-y-4 text-gray-700">
            <li>Share past exam questions with your classmates</li>
            <li>Access a growing archive of exam materials</li>
            <li>Organize questions by course and semester</li>
            <li>Collaborate with fellow students</li>
          </ul>
        </div>
      </div>
    </div>
  );
};

export default BeeArchiveDoc;
