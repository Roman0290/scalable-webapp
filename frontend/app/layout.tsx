import type { Metadata } from "next";
import type { ReactNode } from "react";
import "./globals.css";
import { Footer } from "@/components/layout/footer";
import { Header } from "@/components/layout/header";
import { Providers } from "@/app/providers";

export const metadata: Metadata = {
	title: "SaaS Starter",
	description: "Production-ready full-stack SaaS starter with Next.js and Go.",
};

export default function RootLayout({ children }: { children: ReactNode }) {
	return (
		<html lang="en">
			<body>
				<Providers>
					<div className="min-h-screen">
						<Header />
						<main>{children}</main>
						<Footer />
					</div>
				</Providers>
			</body>
		</html>
	);
}
