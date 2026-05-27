package dashboard

// dashboardHTML contains the embedded UI for the log viewer
const dashboardHTML = `
<!DOCTYPE html>
<html lang="en" data-bs-theme="dark">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Argus Telemetry</title>
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css" rel="stylesheet">
    <style>
        body { background-color: #121212; }
        .log-container { height: 75vh; overflow-y: auto; font-family: monospace; font-size: 0.9rem; }
        .log-VERBOSE { color: #6c757d; }
        .log-WARNING { color: #fd7e14; }
        .log-ERROR { color: #dc3545; }
        .log-CRITICAL { color: #dc3545; font-weight: bold; background-color: rgba(220, 53, 69, 0.1); }
    </style>
</head>
<body>
    <div class="container-fluid py-4">
        <div class="d-flex justify-content-between align-items-center mb-4 border-bottom pb-2">
            <h2 class="text-primary">Argus Telemetry Dashboard</h2>
            <button class="btn btn-outline-light btn-sm" onclick="fetchLogs()">Refresh Logs</button>
        </div>
        
        <div class="card shadow">
            <div class="card-header d-flex justify-content-between align-items-center">
                <h5 class="mb-0">System Logs</h5>
                <span class="badge bg-secondary" id="logCount">0 Entries</span>
            </div>
            <div class="card-body log-container bg-dark" id="logViewer">
                <div class="text-center text-muted mt-5">Loading logs...</div>
            </div>
        </div>
    </div>

    <script>
        async function fetchLogs() {
            try {
                const response = await fetch('/api/logs');
                const logs = await response.json();
                const viewer = document.getElementById('logViewer');
                const logCount = document.getElementById('logCount');
                
                viewer.innerHTML = '';
                logCount.textContent = logs.length + ' Entries';

                if (logs.length === 0) {
                    viewer.innerHTML = '<div class="text-center text-muted mt-5">No logs found.</div>';
                    return;
                }

                logs.forEach(log => {
                    const div = document.createElement('div');
                    div.className = 'log-line log-' + log.level;
                    div.textContent = log.message;
                    viewer.appendChild(div);
                });
                
                // Auto-scroll to bottom
                viewer.scrollTop = viewer.scrollHeight;
            } catch (error) {
                console.error('Failed to fetch logs:', error);
            }
        }

        // Fetch logs on load and set an interval to auto-refresh every 5 seconds
        fetchLogs();
        setInterval(fetchLogs, 5000);
    </script>
</body>
</html>
`
