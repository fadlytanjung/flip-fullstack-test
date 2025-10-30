import { Construction } from 'lucide-react';
import styles from './UnderMaintenance.module.css';

interface UnderMaintenanceProps {
  title?: string;
  message?: string;
}

export function UnderMaintenance({ 
  title = 'Under Maintenance', 
  message = 'This page is currently under development. Please check back later.' 
}: UnderMaintenanceProps) {
  return (
    <div className={styles.container}>
      <div className={styles.card}>
        <Construction className={styles.icon} size={64} />
        <h2 className={styles.title}>{title}</h2>
        <p className={styles.message}>{message}</p>
      </div>
    </div>
  );
}

