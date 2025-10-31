'use client';

import { useState, useEffect, useCallback, useRef } from 'react';
import { Wallet, BarChart3, AlertTriangle, CheckCircle2, XCircle, AlertCircle } from 'lucide-react';
import {
  DataTable,
  Dropzone,
  DashboardLayout,
  CardSkeleton,
} from '@/app/components';
import { useNotification } from '@/app/contexts';
import { useTable, useLoading } from '../hooks';
import { uploadFile, fetchBalance, ApiError, UploadResponse, fetchIssuesCount } from './api';
import styles from './DashboardView.module.css';
import type { DropzoneRef } from '@/app/components/Dropzone/Dropzone';

export function DashboardView() {
  const { showSuccess, showError, showWarning } = useNotification();
  const table = useTable({ perPage: 10 });
  const loading = useLoading();
  const dropzoneRef = useRef<DropzoneRef>(null);

  // Dashboard specific state
  const [balance, setBalance] = useState<any>(null);
  const [issuesCount, setIssuesCount] = useState(0);

  // Fetch balance data
  const fetchBalanceData = useCallback(async () => {
    try {
      const balanceData = await fetchBalance();
      setBalance(balanceData);
    } catch (error) {
      console.error('Failed to fetch balance:', error);
    }
  }, []);

  // Fetch issues count
  const fetchIssuesCountData = useCallback(async () => {
    try {
      const response = await fetchIssuesCount();
      setIssuesCount(response);
    } catch (error) {
      console.error('Failed to fetch issues count:', error);
    }
  }, []);

  // Initial data load
  useEffect(() => {
    loading.setStatsLoading(true);
    Promise.all([
      (async () => {
        await fetchBalanceData();
      })(),
      (async () => {
        await fetchIssuesCountData();
      })()
    ]).then(() => {
      loading.setStatsLoading(false);
    });
  }, [loading.setStatsLoading, fetchBalanceData, fetchIssuesCountData]);

  const handleFileUpload = async (file: File) => {
    // Validate file type
    if (!file.name.endsWith('.csv')) {
      const errorMsg = 'Invalid file type. Please upload a CSV file.';
      showError(errorMsg, 'Upload Failed');
      return;
    }

    try {
      loading.setUploading(true);
      
      const response = await uploadFile(file);

      if ('error' in response) {
        const error = response as ApiError;
        const errorMsg = error.message || 'Failed to upload file';
        showError(errorMsg, 'Upload Failed');
      } else {
        const uploadResp = response as UploadResponse;

        // Show detailed upload results
        const successMsg = `${file.name} uploaded successfully!\n` +
          `Total: ${uploadResp.total_records}, Success: ${uploadResp.success_records}, ` +
          `Failed: ${uploadResp.failed_records}, Pending: ${uploadResp.pending_records}`;

        showSuccess(successMsg, 'Upload Complete');

        // Reset dropzone and refresh data after successful upload
        dropzoneRef.current?.reset();
        setTimeout(() => {
          table.refreshData();
          fetchBalanceData();
          fetchIssuesCountData();
        }, 500);
      }
    } catch (error) {
      const errorMsg = 'Failed to upload file. Please try again.';
      showError(errorMsg, 'Upload Failed');
    } finally {
      loading.setUploading(false);
    }
  };

  // Calculate stats
  const calculateBalance = () => {
    return balance?.balance || 0;
  };

  const getStatusBadge = (status: string) => {
    const statusConfig = {
      SUCCESS: { icon: CheckCircle2, color: styles.statusSuccess },
      FAILED: { icon: XCircle, color: styles.statusFailed },
      PENDING: { icon: AlertCircle, color: styles.statusPending }
    };

    const config = statusConfig[status as keyof typeof statusConfig];
    if (!config) return null;

    const Icon = config.icon;
    return (
      <span className={`${styles.statusBadge} ${config.color}`}>
        <Icon size={14} />
        {status}
      </span>
    );
  };

  // Table columns definition
  const columns = [
    {
      key: 'timestamp',
      header: 'Date',
      sortable: true,
      render: (row: any) => <span className={styles.dateCell}>{table.formatDate(row.timestamp)}</span>
    },
    {
      key: 'name',
      header: 'Name',
      sortable: true,
      render: (row: any) => <span className={styles.nameCell}>{row.name}</span>
    },
    {
      key: 'type',
      header: 'Type',
      sortable: true,
      render: (row: any) => (
        <span className={`${styles.typeBadge} ${row.type === 'CREDIT' ? styles.badgeCredit : styles.badgeDebit}`}>
          {row.type}
        </span>
      )
    },
    {
      key: 'amount',
      header: 'Amount',
      sortable: true,
      render: (row: any) => (
        <span className={`${styles.amountCell} ${row.type === 'CREDIT' ? styles.amountCredit : styles.amountDebit}`}>
          {row.type === 'CREDIT' ? '+' : '-'}{table.formatCurrency(row.amount)}
        </span>
      )
    },
    {
      key: 'status',
      header: 'Status',
      sortable: true,
      render: (row: any) => getStatusBadge(row.status)
    },
    {
      key: 'description',
      header: 'Description',
      render: (row: any) => <span className={styles.descriptionCell}>{row.description}</span>
    }
  ];

  return (
    <DashboardLayout>
      <div className={styles.dashboard}>
        {/* Stats Cards */}
        <div className={styles.statsGrid}>
          {loading.statsLoading ? (
            <>
              <CardSkeleton />
              <CardSkeleton />
              <CardSkeleton />
            </>
          ) : (
            <>
              <div className={styles.statCard}>
                <div className={styles.statIcon}>
                  <Wallet size={32} />
                </div>
                <div className={styles.statContent}>
                  <div className={styles.statLabel}>Current Balance</div>
                    <div className={styles.statValue}>{table.formatCurrency(calculateBalance())}</div>
                </div>
              </div>
              <div className={styles.statCard}>
                <div className={styles.statIcon}>
                  <BarChart3 size={32} />
                </div>
                <div className={styles.statContent}>
                  <div className={styles.statLabel}>Total Transactions</div>
                    <div className={styles.statValue}>{table.apiResponse?.meta.pagination.total || 0}</div>
                </div>
              </div>
              <div className={styles.statCard}>
                <div className={styles.statIcon}>
                  <AlertTriangle size={32} />
                </div>
                <div className={styles.statContent}>
                  <div className={styles.statLabel}>Issues</div>
                    <div className={styles.statValue}>{issuesCount}</div>
                </div>
              </div>
            </>
          )}
        </div>

        {/* Upload Section */}
        <div className={styles.uploadContainer}>
          <Dropzone 
            ref={dropzoneRef}
            onFileSelect={handleFileUpload}
            title="Upload Bank Statement"
            subtitle="Click or drag and drop your CSV file"
            loading={loading.uploading}
          />
        </div>

        {/* All Transactions Table */}
        <DataTable
          title="All Transactions"
          subtitle={table.apiResponse ? `Showing ${table.apiResponse.meta.pagination.count} of ${table.apiResponse.meta.pagination.total} transactions` : undefined}
          columns={columns}
          data={table.transactions}
          meta={table.apiResponse?.meta.pagination}
          loading={table.loading}
          searchValue={table.searchQuery}
          onSearchChange={table.handleSearch}
          onPageChange={table.handlePageChange}
          onSort={table.handleSort}
          sortBy={table.sortBy}
          sortOrder={table.sortOrder}
          showSearch={true}
          showPagination={true}
          emptyMessage="No transactions found. Upload a CSV file to get started."
          filterType={table.filterType}
          onFilterTypeChange={table.handleFilterType}
          filterStatus={table.filterStatus}
          onFilterStatusChange={table.handleFilterStatus}
        />
      </div>
    </DashboardLayout>
  );
}

