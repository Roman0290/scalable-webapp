import { Button } from "@/components/ui/button";
import { Card } from "@/components/ui/card";

const highlights = [
	{
		title: "Auth-first architecture",
		description: "JWT access/refresh flow, protected routes, and a clean auth boundary.",
	},
	{
		title: "Composable frontend",
		description: "App Router, reusable UI primitives, React Query, and Zod-ready forms.",
	},
	{
		title: "Scalable backend",
		description: "Clean Architecture structure for handlers, services, repositories, and persistence.",
	},
];

export default function HomePage() {
	return (
		<section className="mx-auto flex min-h-[calc(100vh-160px)] max-w-7xl flex-col justify-center px-6 py-16 lg:px-10">
			<div className="glass-panel rounded-[2rem] px-8 py-12 shadow-2xl shadow-slate-950/20 lg:px-14 lg:py-16">
				<div className="max-w-3xl space-y-6">
					<span className="inline-flex rounded-full border border-white/10 bg-white/5 px-4 py-1 text-sm text-slate-200">
						Production-ready SaaS foundation
					</span>
					<div className="space-y-4">
						<h1 className="text-4xl font-semibold tracking-tight text-white sm:text-5xl lg:text-6xl">
							Build a secure, scalable full-stack platform with Next.js and Go.
						</h1>
						<p className="max-w-2xl text-base leading-7 text-slate-300 sm:text-lg">
							This starter combines an App Router frontend, a modular Go backend, structured
							authentication boundaries, and a foundation for PostgreSQL, Redis, and CI/CD.
						</p>
					</div>
					<div className="flex flex-col gap-3 sm:flex-row">
						<Button href="/dashboard">Open dashboard</Button>
						<Button href="/login" variant="secondary">
							Sign in
						</Button>
					</div>
				</div>

				<div className="mt-12 grid gap-4 md:grid-cols-3">
					{highlights.map((item) => (
						<Card key={item.title} title={item.title} description={item.description} />
					))}
				</div>
			</div>
		</section>
	);
}
