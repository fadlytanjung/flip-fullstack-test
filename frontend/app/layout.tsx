import type { Metadata } from "next";
import { SidebarProvider, NotificationProvider } from "./contexts";
import "./globals.css";

export const metadata: Metadata = {
  title: "Bank Statement Viewer",
  description: "Upload and view bank statement insights",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body>
        <NotificationProvider>
          <SidebarProvider>
            {children}
          </SidebarProvider>
        </NotificationProvider>
      </body>
    </html>
  );
}
