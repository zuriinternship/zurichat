import Head from 'next/head';

import { PluginContent } from '../components/PluginContent';
import { Sidebar } from '../components/Sidebar';
import { Topbar } from '../components/Topbar';
import styles from '../styles/Home.module.css';

export default function Home() {
  return (
    <div className={styles.container}>
      <Head>
        <title>Zuri Chat</title>
        <link rel="icon" href="/zurichatlogo.svg" />
      </Head>
      <Sidebar />
      <div className={styles.room}>
        <Topbar />
        <div className={styles.pluginContent}>
          <PluginContent />
        </div>
      </div>
    </div>
  );
}
