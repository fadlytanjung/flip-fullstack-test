'use client';

import { useState } from 'react';
import { DashboardLayout, Alert, Button, Pagination } from '@/app/components';
import { useNotification } from '@/app/contexts';
import styles from './ComponentDemo.module.css';

export function ComponentDemo() {
  const { showSuccess, showError, showWarning, showInfo } = useNotification();
  const [currentPage, setCurrentPage] = useState(5);
  const [showAlert, setShowAlert] = useState(true);

  // Mock pagination data for demo
  const paginationMeta = {
    total: 100,
    count: 10,
    per_page: 10,
    current_page: currentPage,
    total_pages: 10,
    links: {
      next: currentPage < 10 ? `/demo?page=${currentPage + 1}` : undefined,
      prev: currentPage > 1 ? `/demo?page=${currentPage - 1}` : undefined,
    }
  };

  return (
    <DashboardLayout>
      <div className={styles.demoContainer}>
        <h1 className={styles.title}>Component Demo</h1>
        <p className={styles.subtitle}>Testing Alert, Notification, and Smart Pagination</p>

        {/* Alert Demo Section */}
        <section className={styles.section}>
          <h2 className={styles.sectionTitle}>Alert Components</h2>
          <p className={styles.sectionDescription}>
            Static inline alerts for persistent feedback
          </p>
          
          <div className={styles.alertGrid}>
            <Alert
              type="success"
              title="Success!"
              message="Your operation completed successfully."
            />
            
            <Alert
              type="error"
              title="Error!"
              message="Something went wrong. Please try again."
            />
            
            <Alert
              type="warning"
              title="Warning!"
              message="This action may have unintended consequences."
            />
            
            <Alert
              type="info"
              title="Info"
              message="Here's some helpful information for you."
            />

            {showAlert && (
              <Alert
                type="success"
                title="Dismissible Alert"
                message="This alert can be closed by clicking the X button."
                dismissible
                onClose={() => setShowAlert(false)}
              />
            )}
          </div>

          {!showAlert && (
            <Button onClick={() => setShowAlert(true)}>
              Show Dismissible Alert
            </Button>
          )}
        </section>

        {/* Notification Demo Section */}
        <section className={styles.section}>
          <h2 className={styles.sectionTitle}>Notification (Toast) Components</h2>
          <p className={styles.sectionDescription}>
            Temporary toast notifications that auto-dismiss
          </p>
          
          <div className={styles.buttonGroup}>
            <Button 
              variant="primary"
              onClick={() => showSuccess('This is a success notification!')}
            >
              Show Success
            </Button>
            
            <Button 
              variant="secondary"
              onClick={() => showError('This is an error notification!', 'Error Title')}
            >
              Show Error
            </Button>
            
            <Button 
              variant="outline"
              onClick={() => showWarning('This is a warning notification!', 'Warning')}
            >
              Show Warning
            </Button>
            
            <Button 
              variant="outline"
              onClick={() => showInfo('This is an info notification!', 'Info')}
            >
              Show Info
            </Button>
          </div>

          <div className={styles.demoCode}>
            <code>
              {`// Usage\nconst { showSuccess, showError } = useNotification();\n\nshowSuccess('Operation completed!');\nshowError('Something went wrong', 'Error');`}
            </code>
          </div>
        </section>

        {/* Pagination Demo Section */}
        <section className={styles.section}>
          <h2 className={styles.sectionTitle}>Smart Pagination</h2>
          <p className={styles.sectionDescription}>
            Intelligent ellipsis display based on current page
          </p>

          <div className={styles.paginationDemo}>
            <div className={styles.paginationInfo}>
              <p><strong>Current Page:</strong> {currentPage} of 10</p>
              <p><strong>Pattern:</strong></p>
              <ul className={styles.patternList}>
                <li>Pages 1-3: <code>[1][2][3][4][5]...[10]</code></li>
                <li>Page 5: <code>[1]...[4][5][6]...[10]</code></li>
                <li>Pages 8-10: <code>[1]...[6][7][8][9][10]</code></li>
              </ul>
            </div>

            <Pagination
              meta={paginationMeta}
              onPageChange={(page) => {
                setCurrentPage(page);
                showInfo(`Navigated to page ${page}`, 'Page Changed');
              }}
            />
          </div>

          <div className={styles.quickJumpButtons}>
            <p><strong>Quick Jump:</strong></p>
            <div className={styles.buttonGroup}>
              <Button size="small" onClick={() => setCurrentPage(1)}>Page 1</Button>
              <Button size="small" onClick={() => setCurrentPage(3)}>Page 3</Button>
              <Button size="small" onClick={() => setCurrentPage(5)}>Page 5</Button>
              <Button size="small" onClick={() => setCurrentPage(8)}>Page 8</Button>
              <Button size="small" onClick={() => setCurrentPage(10)}>Page 10</Button>
            </div>
          </div>
        </section>

        {/* Combined Demo Section */}
        <section className={styles.section}>
          <h2 className={styles.sectionTitle}>Real-World Example</h2>
          <p className={styles.sectionDescription}>
            Simulating form submission with alerts and notifications
          </p>

          <div className={styles.exampleCard}>
            <Alert
              type="info"
              message="Fill out this demo form and submit to see the components in action."
            />

            <form 
              className={styles.demoForm}
              onSubmit={(e) => {
                e.preventDefault();
                showWarning('Processing your request...', 'Submitting');
                
                setTimeout(() => {
                  showSuccess('Form submitted successfully!', 'Success');
                }, 1500);
              }}
            >
              <input 
                type="text" 
                placeholder="Enter your name" 
                className={styles.input}
                required
              />
              <input 
                type="email" 
                placeholder="Enter your email" 
                className={styles.input}
                required
              />
              <Button type="submit" variant="primary">
                Submit Demo Form
              </Button>
            </form>
          </div>
        </section>
      </div>
    </DashboardLayout>
  );
}

