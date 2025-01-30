import React from 'react';

const WhatsNew: React.FC = () => {
    return (
        <div className="min-h-screen bg-white py-16 px-4 sm:px-6 lg:px-8">
            <div className="max-w-4xl mx-auto">
                <h1 className="text-4xl font-bold text-[#212121] mb-8">What's New in BeeHub</h1>

                {/* Latest Version Banner */}
                <div className="bg-[#F5FDFD] p-6 rounded-lg shadow-md mb-12">
                    <div className="flex items-center gap-2 mb-4">
                        <span className="text-sm bg-[#FDC003] text-[#212121] px-3 py-1 rounded-full">Latest Update</span>
                        <span className="text-gray-500">Version 1.1.0</span>
                    </div>

                    {/* Major Changes Section */}
                    <div className="mb-8">
                        <h2 className="text-2xl font-semibold text-[#0372CE] mb-4">Major Changes</h2>
                        <div className="space-y-4">
                            {/* Open Source */}
                            <div className="bg-white p-4 rounded-lg border border-gray-100">
                                <h3 className="font-semibold text-[#212121] mb-2">🌟 Now Open Source</h3>
                                <p className="text-gray-600">
                                    BeeHub is now open source! Join our community and contribute to making course registration better for everyone.
                                </p>
                            </div>

                            {/* Reserve CRN Feature */}
                            <div className="bg-white p-4 rounded-lg border border-gray-100">
                                <h3 className="font-semibold text-[#212121] mb-2">🔄 Reserve Course System</h3>
                                <p className="text-gray-600 mb-3">
                                    Introducing the new Reserve Course feature! Now you can set backup courses in case your primary choice becomes unavailable.
                                </p>
                                <div className="bg-[#F5FDFD] p-3 rounded-md">
                                    <p className="text-sm text-gray-600">
                                        <span className="font-medium">How it works:</span><br />
                                        1. Select your primary course<br />
                                        2. Add alternative courses as reserves<br />
                                        3. If the primary course is full, BeeHub will automatically try to register for your reserve choices
                                    </p>
                                </div>
                            </div>
                        </div>
                    </div>

                    {/* Feature Improvements Section */}
                    <h2 className="text-2xl font-semibold text-[#0372CE] mb-4">Feature Improvements</h2>
                    <ul className="space-y-3">
                        <li className="flex items-start gap-3">
                            <span className="text-[#FDC003] mt-1">•</span>
                            <div>
                                <h3 className="font-medium text-[#212121] mb-1">Direct CRN Input</h3>
                                <p className="text-gray-600">
                                    Quickly add courses by directly entering CRN numbers, making the course selection process faster and more efficient.
                                </p>
                            </div>
                        </li>

                        <li className="flex items-start gap-3">
                            <span className="text-[#FDC003] mt-1">•</span>
                            <div>
                                <h3 className="font-medium text-[#212121] mb-1">Auto-Update System</h3>
                                <p className="text-gray-600">
                                    BeeHub now automatically checks for and installs updates, ensuring you always have the latest features and improvements.
                                </p>
                            </div>
                        </li>

                        <li className="flex items-start gap-3">
                            <span className="text-[#FDC003] mt-1">•</span>
                            <div>
                                <h3 className="font-medium text-[#212121] mb-1">Transfer Student Login Fix</h3>
                                <p className="text-gray-600">
                                    Resolved login issues for transfer students, ensuring smooth access to the platform for all users.
                                </p>
                            </div>
                        </li>
                    </ul>

                    {/* Interface Updates Section */}
                    <h2 className="text-2xl font-semibold text-[#0372CE] mt-8 mb-4">Interface Updates</h2>
                    <div className="bg-white p-4 rounded-lg border border-gray-100">
                        <p className="text-gray-600">
                            We've streamlined the interface by removing unused application pages and focusing on core features that matter most to our users.
                        </p>
                    </div>
                </div>
            </div>
        </div>
    );
};

export default WhatsNew;
