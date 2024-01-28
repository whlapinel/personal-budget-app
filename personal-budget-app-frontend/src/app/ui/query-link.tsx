'use client';

import {Link} from "@/app/ui/link";
import { usePathname } from "next/navigation";

export default function QueryLink ({query, children}) {
    const pathname = usePathname();
    return (
        <Link
            href={{
                pathname: pathname,
                query,
            }}
        >
            {children}
        </Link>
    )
}