import React from 'react';

const BeePickerPage: React.FC = () => {
  return (
    <div className="min-h-screen flex flex-col bg-[#F5FDFD] h-screen overflow-y-auto">
      {/* Scrollable container */}
      <div className="flex-1 overflow-y-scroll px-4 py-8">
        <div className="max-w-4xl w-full mx-auto p-8 bg-white rounded-lg shadow-lg">

          <h1 className="text-4xl font-bold text-[#0372CE] text-center mb-6">
            BeePicker
          </h1>
          <p className="text-lg text-gray-600 text-center mb-10">
            BeePicker offers a streamlined solution to manage your course selections. Avoid conflicts, and automize your perfect schedule with ease.
          </p>

          {/* Screenshot Image */}
          <div className="bg-gray-100 rounded-lg p-6 mb-6">
            <img
              src="/images/beepicker_screenshot.png"  // Adjust the path based on where you store the screenshot
              alt="BeePicker Screenshot"
              className="rounded"
            />
          </div>

          <div className="text-left space-y-6">
            <div>
              <h2 className="text-2xl font-semibold text-gray-800 mb-2">Step 1: Browse Available Courses</h2>
              <p className="text-gray-600">
                Navigate through the list of available courses. Use filters based on faculties, departments, and course levels to narrow down your selection.
              </p>
            </div>

            <div>
              <h2 className="text-2xl font-semibold text-gray-800 mb-2">Step 2: Select Courses</h2>
              <p className="text-gray-600">
                Once you find courses you are interested in, add them to your list. BeePicker will notify you of any schedule conflicts.
              </p>
            </div>

            <div>
              <h2 className="text-2xl font-semibold text-gray-800 mb-2">Step 3: Submit Course Selection</h2>
              <p className="text-gray-600">
                During the course selection period, click the "Submit Course Selection" button. This will attempt to pick your desired courses within a short time frame. After submission, the results of your course selections will be displayed on the page.
              </p>
            </div>

          </div>
        </div>
        {/* Added these so that footer isnt colliding with content*/}
        <br />
        <br />
        <br />
        <br />
      </div>
    </div>

  );
};

export default BeePickerPage;
