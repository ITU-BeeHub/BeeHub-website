import { Link } from "react-router-dom";

export default function Footer() {
  return (
    <footer className="absolute bottom-4 left-1/2 transform -translate-x-1/2 flex space-x-10 text-[#0372CE] z-10">
      <Link className="block" to={'/documentation'}>
        Documentation
      </Link>
      <Link className="block" to={'/contact'}>
        Contact Us
      </Link>
    </footer>
  );
}
