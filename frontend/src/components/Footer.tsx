import { Link } from "react-router-dom";

export default function Footer() {
  return (
    <footer className="py-2 mt-auto bg-white border-t">
      <div className="container mx-auto flex justify-center space-x-10 text-[#0372CE]">
        <Link className="hover:text-[#0372CEdd] transition-colors" to={'/documentation'}>
          Documentation
        </Link>
        <Link className="hover:text-[#0372CEdd] transition-colors" to={'/contact'}>
          Contact Us
        </Link>
      </div>
    </footer>
  );
}
