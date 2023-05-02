import '@fortawesome/fontawesome-svg-core/styles.css';
import { library } from '@fortawesome/fontawesome-svg-core';
import { faDollar, faUser } from '@fortawesome/free-solid-svg-icons';
import '@/styles/globals.css'
import type { AppProps } from 'next/app'

library.add(faDollar, faUser);

export default function App({ Component, pageProps }: AppProps) {
  return <Component {...pageProps} />
}
