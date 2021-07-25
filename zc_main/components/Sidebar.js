import React from 'react';

import { channels } from '../lib/channels';
import { messages } from '../lib/dms';
import styles from '../styles/Sidebar.module.css';

export const Sidebar = () => {
  return (
    <div className={styles.container}>
      <div className={styles.zuriLogo}>
        <img src="/zurichatlogo.svg" alt="Zuri Chat logo" />
        <p>ZURI</p>
      </div>
      <div className={styles.orgInfo}>
        <div className={styles.orgName}>
          <p>HNGi8</p>
          <img
            src="/shapekeyboardarrowdown.svg"
            alt="Organisation settings button"
          />
        </div>
        <div className={styles.newMessage}>
          <img src="/newmessage.svg" alt="New message icon" />
        </div>
      </div>
      <div className={styles.channelsContainer}>
        <div className={styles.channelHeader}>
          <img
            className={styles.channelDropdownIcon}
            src="/shapekeyboardarrowdown2.svg"
            alt="Channels dropdown button"
          />
          <p>Channels</p>
          <img
            className={styles.channelAddBtn}
            src="/addicon.svg"
            alt="Add Channels button"
          />
        </div>
        <div className={styles.channelNames}>
          {channels.map(channel => (
            <div key={channel.id}>
              <p>
                <span>#</span>
                {channel.name}
              </p>
            </div>
          ))}
        </div>
      </div>
      <div className={styles.messagesContainer}>
        <div className={styles.messageHeader}>
          <img
            className={styles.messageDropdownIcon}
            src="/shapekeyboardarrowdown2.svg"
            alt="messages dropdown button"
          />
          <p>messages</p>
          <img
            className={styles.messageAddBtn}
            src="/addicon.svg"
            alt="Add messages button"
          />
        </div>
        <div className={styles.messageNames}>
          {messages.map(message => (
            <div className={styles.messageTitle} key={message.id}>
              <p>
                <span>
                  <img src={message.avatar} alt="avatar" />
                </span>
                {message.name}
              </p>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
};
