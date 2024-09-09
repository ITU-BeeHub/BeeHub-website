import { Link, useNavigate } from "react-router-dom";
import BeakerIcon from "./icons/BeakerIcon";

export default function Header() {

  return (
    <header className="flex h-[60px] items-center justify-between border-b bg-white px-4 lg:px-6">
      <Link to="/" className="flex items-center gap-2">
        <BeakerIcon className="h-8 w-8 text-[#FDC003]" />
        <span className="text-2xl font-bold text-[#212121]">BeeHub</span>
      </Link>
    </header>
  );
}
