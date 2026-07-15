export default function DashboardPage() {
  return (
    <section className="mx-auto max-w-7xl px-6 py-16 lg:px-10">
      <div className="glass-panel rounded-[2rem] p-8 lg:p-12">
        <p className="text-sm uppercase tracking-[0.25em] text-slate-400">Dashboard</p>
        <h1 className="mt-3 text-3xl font-semibold text-white">Protected workspace</h1>
        <p className="mt-4 max-w-2xl text-slate-300">
          This route is ready for auth protection, role-based access, and API-driven data views.
        </p>
      </div>
    </section>
  );
}