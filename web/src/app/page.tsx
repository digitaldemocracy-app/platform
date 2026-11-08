import { appInfo } from '@/lib/app-info';

const links = [
  { label: 'Manifesto', href: 'https://github.com/digitaldemocracy-app/manifesto' },
  { label: 'Docs & roadmap', href: 'https://github.com/digitaldemocracy-app/docs' },
  { label: 'Source code', href: 'https://github.com/digitaldemocracy-app/platform' },
];

export default function Home() {
  const { name, version } = appInfo();
  return (
    <main className="mx-auto flex min-h-screen max-w-2xl flex-col justify-center gap-6 px-4 py-12 sm:px-6">
      <h1 className="text-4xl font-bold tracking-tight sm:text-5xl">Digital Democracy</h1>
      <p className="text-lg text-slate-700">
        The platform is being built in public. Nothing here yet beyond this page.
      </p>
      <nav aria-label="Project links">
        <ul className="flex flex-col gap-2">
          {links.map((l) => (
            <li key={l.href}>
              <a
                href={l.href}
                className="text-blue-700 underline underline-offset-4 hover:text-blue-900 focus-visible:outline-2 focus-visible:outline-offset-2"
              >
                {l.label}
              </a>
            </li>
          ))}
        </ul>
      </nav>
      <p className="text-sm text-slate-600">
        {name} v{version}
      </p>
    </main>
  );
}
