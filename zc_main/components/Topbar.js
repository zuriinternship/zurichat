import React from 'react';

import styles from '../styles/Topbar.module.css';

export const Topbar = () => {
  return (
    <div className={styles.container}>
      <img src="/settings.svg" alt="settings" />
      <div className={styles.profile}>
        <img src="/profilepic.png" alt="Profile" />
      </div>
    </div>
  );
};
