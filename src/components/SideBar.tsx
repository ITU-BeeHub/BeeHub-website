import { NavLink } from "react-router-dom";
import { useState, useEffect } from "react";

import ArchiveIcon from "./icons/ArchiveIcon";
import CalendarIcon from "./icons/CalendarIcon";
import FolderSyncIcon from "./icons/FolderSyncIcon";
import WebcamIcon from "./icons/WebCamIcon";
import PickaxeIcon from "./icons/PickaxeIcon";
import HamburgerIcon from "./icons/Hamburger";

export default function Sidebar() {
  const [isCollapsed, setIsCollapsed] = useState(false);

  // Collapse sidebar on narrow screens
  useEffect(() => {
    const mediaQuery = window.matchMedia("(max-width: 768px)");

    const handleMediaQueryChange = (event: MediaQueryListEvent) => {
      setIsCollapsed(event.matches);
    };

    // Set initial state
    setIsCollapsed(mediaQuery.matches);

    // Add event listener
    mediaQuery.addEventListener("change", handleMediaQueryChange);

    // Clean up event listener
    return () => {
      mediaQuery.removeEventListener("change", handleMediaQueryChange);
    };
  }, []);

  const toggleSidebar = () => {
    setIsCollapsed(!isCollapsed);
  };

  return (
    <div className="flex h-screen">
      {/* Sidebar */}
      <aside
        className={`bg-white h-full border-r p-4 transition-all duration-300 flex flex-col justify-between ${
          isCollapsed ? "w-20" : "w-[200px]"
        }`}
      >
        <div>
          <nav className="px-2">
            <div
              className="flex items-center gap-2 py-4 cursor-pointer"
              onClick={toggleSidebar}
            >
              <HamburgerIcon className="h-8 w-8 flex-shrink-0 text-[#FDC003]" />
              <span
                className={`overflow-hidden transition-transform duration-300 ${
                  isCollapsed ? "opacity-0 scale-0" : "opacity-100 scale-100"
                }`}
              >
                <b>BeeHub</b>
              </span>
            </div>
          </nav>
          <nav className="mt-8 flex flex-col gap-2">
            <NavLink
              to="/beepicker"
              className={({ isActive }) =>
                `flex items-center gap-3 p-3 transition-all duration-0 ${
                  isActive
                    ? "border-l-4 border-[#FDC003] bg-[#F5FDFD] text-[#212121] -mr-4 rounded-l-lg"
                    : "text-[#212121] hover:bg-[#FDC003] hover:text-white rounded-lg"
                }`
              }
            >
              <PickaxeIcon className="h-6 w-6 flex-shrink-0" />
              <span
                className={`overflow-hidden transition-transform duration-300 ${
                  isCollapsed ? "opacity-0 scale-0" : "opacity-100 scale-100"
                }`}
              >
                BeePicker
              </span>
            </NavLink>
            <NavLink
              to="/beesync"
              className={({ isActive }) =>
                `flex items-center gap-3 p-3 transition-all duration-0 ${
                  isActive
                    ? "border-l-4 border-[#FDC003] bg-[#F5FDFD] text-[#212121] -mr-4 rounded-l-lg"
                    : "text-[#212121] hover:bg-[#FDC003] hover:text-white rounded-lg"
                }`
              }
            >
              <FolderSyncIcon className="h-6 w-6 flex-shrink-0" />
              <span
                className={`overflow-hidden transition-transform duration-300 ${
                  isCollapsed ? "opacity-0 scale-0" : "opacity-100 scale-100"
                }`}
              >
                BeeSync
              </span>
            </NavLink>
            <NavLink
              to="/beecalendar"
              className={({ isActive }) =>
                `flex items-center gap-3 p-3 transition-all duration-0 ${
                  isActive
                    ? "border-l-4 border-[#FDC003] bg-[#F5FDFD] text-[#212121] -mr-4 rounded-l-lg"
                    : "text-[#212121] hover:bg-[#FDC003] hover:text-white rounded-lg"
                }`
              }
            >
              <CalendarIcon className="h-6 w-6 flex-shrink-0" />
              <span
                className={`overflow-hidden transition-transform duration-300 ${
                  isCollapsed ? "opacity-0 scale-0" : "opacity-100 scale-100"
                }`}
              >
                BeeCalendar
              </span>
            </NavLink>
            <NavLink
              to="/beearchive"
              className={({ isActive }) =>
                `flex items-center gap-3 p-3 transition-all duration-0 ${
                  isActive
                    ? "border-l-4 border-[#FDC003] bg-[#F5FDFD] text-[#212121] -mr-4 rounded-l-lg"
                    : "text-[#212121] hover:bg-[#FDC003] hover:text-white rounded-lg"
                }`
              }
            >
              <ArchiveIcon className="h-6 w-6 flex-shrink-0" />
              <span
                className={`overflow-hidden transition-transform duration-300 ${
                  isCollapsed ? "opacity-0 scale-0" : "opacity-100 scale-100"
                }`}
              >
                BeeArchive
              </span>
            </NavLink>
            <NavLink
              to="/beechat"
              className={({ isActive }) =>
                `flex items-center gap-3 p-3 transition-all duration-0 ${
                  isActive
                    ? "border-l-4 border-[#FDC003] bg-[#F5FDFD] text-[#212121] -mr-4 rounded-l-lg"
                    : "text-[#212121] hover:bg-[#FDC003] hover:text-white rounded-lg"
                }`
              }
            >
              <WebcamIcon className="h-6 w-6 flex-shrink-0" />
              <span
                className={`overflow-hidden transition-transform duration-300 ${
                  isCollapsed ? "opacity-0 scale-0" : "opacity-100 scale-100"
                }`}
              >
                BeeChat
              </span>
            </NavLink>
          </nav>
        </div>

      </aside>

    </div>
  );
}
