import type { Metadata } from 'next';
import './globals.css';

const description =
  'An open, global platform for constructive debate, secret verifiable voting and democratic participation. Built in public.';

export const metadata: Metadata = {
  title: 'Digital Democracy',
  description,
  metadataBase: new URL('https://digitaldemocracy.app'),
  openGraph: {
    title: 'Digital Democracy',
    description,
    siteName: 'Digital Democracy',
    type: 'website',
  },
  twitter: {
    card: 'summary_large_image',
    title: 'Digital Democracy',
    description,
  },
};

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="en">
      <body className="bg-white text-slate-900 antialiased">{children}</body>
    </html>
  );
}
