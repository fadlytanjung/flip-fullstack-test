'use client';

import { useState, useRef } from 'react';
import { Upload, FileText, X } from 'lucide-react';
import styles from './Dropzone.module.css';

interface DropzoneProps {
  onFileSelect?: (file: File) => void;
  accept?: string;
  title?: string;
  subtitle?: string;
}

export function Dropzone({
  onFileSelect,
  accept = '.csv',
  title = 'Upload File',
  subtitle = 'Click or drag and drop'
}: DropzoneProps) {
  const [isDragging, setIsDragging] = useState(false);
  const [selectedFile, setSelectedFile] = useState<File | null>(null);
  const fileInputRef = useRef<HTMLInputElement>(null);

  const handleDragOver = (e: React.DragEvent) => {
    e.preventDefault();
    setIsDragging(true);
  };

  const handleDragLeave = () => {
    setIsDragging(false);
  };

  const handleDrop = (e: React.DragEvent) => {
    e.preventDefault();
    setIsDragging(false);

    const files = e.dataTransfer.files;
    if (files.length > 0) {
      handleFileSelect(files[0]);
    }
  };

  const handleClick = () => {
    fileInputRef.current?.click();
  };

  const handleFileInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const files = e.target.files;
    if (files && files.length > 0) {
      handleFileSelect(files[0]);
    }
  };

  const handleFileSelect = (file: File) => {
    const fileExtension = '.' + file.name.split('.').pop()?.toLowerCase();
    const acceptedExtensions = accept.split(',').map(ext => ext.trim().toLowerCase());

    if (acceptedExtensions.includes(fileExtension)) {
      setSelectedFile(file);
      if (onFileSelect) {
        onFileSelect(file);
      }
    } else {
      alert(`Please select a valid file type: ${accept}`);
    }
  };

  const handleRemoveFile = (e: React.MouseEvent) => {
    e.stopPropagation();
    setSelectedFile(null);
    if (fileInputRef.current) {
      fileInputRef.current.value = '';
    }
  };

  const formatFileSize = (bytes: number) => {
    if (bytes === 0) return '0 Bytes';
    const k = 1024;
    const sizes = ['Bytes', 'KB', 'MB', 'GB'];
    const i = Math.floor(Math.log(bytes) / Math.log(k));
    return Math.round(bytes / Math.pow(k, i) * 100) / 100 + ' ' + sizes[i];
  };

  return (
    <div className={styles.dropzone}>
      <input
        ref={fileInputRef}
        type="file"
        accept={accept}
        onChange={handleFileInputChange}
        className={styles.hiddenInput}
      />
      
      <div
        className={`${styles.dropArea} ${isDragging ? styles.dragging : ''} ${selectedFile ? styles.hasFile : ''}`}
        onDragOver={handleDragOver}
        onDragLeave={handleDragLeave}
        onDrop={handleDrop}
        onClick={!selectedFile ? handleClick : undefined}
      >
        {!selectedFile ? (
          <div className={styles.emptyState}>
            <Upload className={styles.uploadIcon} size={32} />
            <div className={styles.text}>
              <p className={styles.title}>{title}</p>
              <p className={styles.subtitle}>{subtitle}</p>
            </div>
            <p className={styles.hint}>{accept.replace(/\./g, '').toUpperCase()} only</p>
          </div>
        ) : (
          <div className={styles.filePreview}>
            <FileText className={styles.fileIcon} size={32} />
            <div className={styles.fileInfo}>
              <p className={styles.fileName}>{selectedFile.name}</p>
              <p className={styles.fileSize}>{formatFileSize(selectedFile.size)}</p>
            </div>
            <button
              className={styles.removeButton}
              onClick={handleRemoveFile}
              type="button"
              aria-label="Remove file"
            >
              <X size={18} />
            </button>
          </div>
        )}
      </div>
    </div>
  );
}

