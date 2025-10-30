import { DashboardLayout } from '@/app/components';
import { UnderMaintenance } from '@/app/ui/UnderMaintenance';

export default function UploadPage() {
  return (
    <DashboardLayout>
      <UnderMaintenance 
        title="Upload Page"
        message="The upload page is currently under development. Please check back later."
      />
    </DashboardLayout>
  );
}

