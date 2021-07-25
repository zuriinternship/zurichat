import axios from 'axios';
import React, { useEffect } from 'react';

import styles from '../styles/PluginContent.module.css';

// This component would fetch and display the HTML for the plugin.
// Data about the plugin should be gotten from what ever state management we choose to use.
// The data should contain the URL to call for the plugin data, and that should replace
// the variable used to store the URL currently.
export const PluginContent = () => {
  const pluginUrl = '/apps/default';
  useEffect(() => {
    axios.get(pluginUrl).then(response => {
      const parser = new DOMParser();
      const html = parser.parseFromString(response.data, 'text/html');
      document.querySelector('.plugin-content').innerHTML = html.body.innerHTML;
    });
  }, []);

  return (
    <section className={styles.container}>
      <div className="plugin-content">Loading...</div>
    </section>
  );
};
