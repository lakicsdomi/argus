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
        :root {
            --bg-darker: #0f172a;
            --bg-card: #1e293b;
            --accent: #f59e0b; /* Amber */
            --border-color: #334155;
        }
        body { background-color: var(--bg-darker); color: #cbd5e1; font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif; }
        
        /* Layout & Cards */
        .header-title { color: var(--accent); font-weight: 600; letter-spacing: 0.5px; }
        .card { background-color: var(--bg-card); border: 1px solid var(--border-color); border-radius: 8px; overflow: hidden; }
        .card-header { border-bottom: 1px solid var(--border-color); background-color: rgba(0,0,0,0.2); padding: 16px 20px; }
        .btn-accent { border-color: var(--accent); color: var(--accent); transition: all 0.2s; }
        .btn-accent:hover { background-color: var(--accent); color: #000; }

        /* Log Table Styling */
        .log-container { height: 75vh; overflow-y: auto; padding: 0; background-color: #0b1120; }
        .log-table { width: 100%; border-collapse: collapse; font-family: 'Consolas', 'Monaco', 'Courier New', monospace; font-size: 0.85rem; }
        .log-table th { 
            position: sticky; top: 0; background-color: var(--bg-card); 
            padding: 10px 16px; border-bottom: 1px solid var(--border-color); 
            text-align: left; z-index: 1; font-size: 0.75rem; text-transform: uppercase; 
            letter-spacing: 1px; color: #64748b; 
        }
        .log-row { border-bottom: 1px solid #1e293b; transition: background-color 0.15s; }
        .log-row:hover { background-color: #1e293b; }
        .log-cell { padding: 10px 16px; vertical-align: top; }

        /* Column specific styling */
        .log-time { color: #94a3b8; white-space: nowrap; width: 180px; }
        .log-comp { color: #38bdf8; font-weight: 600; white-space: nowrap; width: 150px; }
        .log-msg { color: #e2e8f0; word-break: break-word; }

        /* Level Badges */
        .badge { padding: 5px 8px; font-weight: 600; letter-spacing: 0.5px; width: 80px; text-align: center; }
        .badge-VERBOSE { background-color: #334155; color: #cbd5e1; }
        .badge-WARNING { background-color: #b45309; color: #fde68a; }
        .badge-ERROR { background-color: #991b1b; color: #fecaca; }
        .badge-CRITICAL { background-color: #7f1d1d; color: #fca5a5; font-weight: bold; border: 1px solid #ef4444; }
    </style>
</head>
<body>
    <div class="container-fluid py-4 px-4">
        <div class="d-flex justify-content-between align-items-center mb-4">
            <h3 class="header-title mb-0">Argus Telemetry</h3>
            
            <div class="d-flex gap-3 align-items-center">
                <select id="levelFilter" class="form-select form-select-sm bg-dark text-light border-secondary" style="width: 140px;" onchange="applyFilter()">
                    <option value="ALL">All Levels</option>
                    <option value="VERBOSE">VERBOSE</option>
                    <option value="WARNING">WARNING</option>
                    <option value="ERROR">ERROR</option>
                    <option value="CRITICAL">CRITICAL</option>
                </select>
                
                <button class="btn btn-sm btn-accent d-flex align-items-center gap-1" onclick="fetchLogs()">
                    <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" fill="currentColor" class="bi bi-arrow-clockwise" viewBox="0 0 16 16">
                        <path fill-rule="evenodd" d="M8 3a5 5 0 1 0 4.546 2.914.5.5 0 0 1 .908-.417A6 6 0 1 1 8 2v1z"/>
                        <path d="M8 4.466V.534a.25.25 0 0 1 .41-.192l2.36 1.966c.12.1.12.284 0 .384L8.41 4.658A.25.25 0 0 1 8 4.466z"/>
                    </svg>
                    Refresh
                </button>
            </div>
        </div>
        
        <div class="card shadow-lg">
            <div class="card-header d-flex justify-content-between align-items-center">
                <h6 class="mb-0 text-light">System Logs</h6>
                <span class="badge bg-secondary" id="logCount">0 Entries</span>
            </div>
            <div class="card-body log-container">
                <table class="log-table">
                    <thead>
                        <tr>
                            <th>Timestamp</th>
                            <th>Level</th>
                            <th>Component</th>
                            <th>Message</th>
                        </tr>
                    </thead>
                    <tbody id="logViewer">
                        <tr>
                            <td colspan="4" class="text-center text-muted py-5">Loading telemetry data...</td>
                        </tr>
                    </tbody>
                </table>
            </div>
        </div>
    </div>

    <script>
        let allParsedLogs = []; // Global store for client-side filtering

        async function fetchLogs() {
            try {
                const response = await fetch('/api/logs');
                const logs = await response.json();
                
                if (!logs || logs.length === 0) {
                    document.getElementById('logViewer').innerHTML = '<tr><td colspan="4" class="text-center text-muted py-5">No logs found in directory.</td></tr>';
                    document.getElementById('logCount').textContent = '0 Entries';
                    return;
                }

                // Parse and format logs
                allParsedLogs = logs.map(log => {
                    const regex = /^\[(.*?)\] (.*?): (.*?): (.*)$/;
                    const match = log.message.match(regex);
                    
                    if (match) {
                        return { timestamp: match[1], level: log.level, component: match[3], msg: match[4], parsed: true, raw: log.message };
                    }
                    
                    // Fallback for non-standard formats: attempt to extract timestamp
                    const timeMatch = log.message.match(/^\[(.*?)\]/);
                    const tstamp = timeMatch ? timeMatch[1] : '1970-01-01 00:00:00';
                    return { timestamp: tstamp, level: log.level, parsed: false, raw: log.message };
                });

                // STRICT CHRONOLOGICAL SORT based on timestamp string
                allParsedLogs.sort((a, b) => a.timestamp.localeCompare(b.timestamp));

                renderLogs();
            } catch (error) {
                console.error('Failed to fetch logs:', error);
            }
        }

        function renderLogs() {
            const tbody = document.getElementById('logViewer');
            const filter = document.getElementById('levelFilter').value;
            tbody.innerHTML = '';

            let visibleCount = 0;

            allParsedLogs.forEach(log => {
                // Apply the selected filter
                if (filter !== 'ALL' && log.level !== filter) return;

                visibleCount++;
                const tr = document.createElement('tr');
                tr.className = 'log-row';

                if (log.parsed) {
                    // Replaced template literals with string concatenation to avoid Go backtick collision
                    tr.innerHTML = '<td class="log-cell log-time">' + log.timestamp + '</td>' +
                                   '<td class="log-cell"><span class="badge badge-' + log.level + '">' + log.level + '</span></td>' +
                                   '<td class="log-cell log-comp">' + log.component + '</td>' +
                                   '<td class="log-cell log-msg">' + log.msg + '</td>';
                } else {
                    const timeStr = log.timestamp !== '1970-01-01 00:00:00' ? log.timestamp : '-';
                    tr.innerHTML = '<td class="log-cell log-time">' + timeStr + '</td>' +
                                   '<td class="log-cell"><span class="badge badge-' + log.level + '">' + log.level + '</span></td>' +
                                   '<td class="log-cell log-msg" colspan="2">' + log.raw + '</td>';
                }
                tbody.appendChild(tr);
            });
            
            document.getElementById('logCount').textContent = visibleCount + ' Entries';

            // Auto-scroll to bottom
            const container = document.querySelector('.log-container');
            container.scrollTop = container.scrollHeight;
        }

        // Triggered by the dropdown change event
        function applyFilter() {
            renderLogs();
        }

        // Fetch logs on load and set auto-refresh
        fetchLogs();
        setInterval(fetchLogs, 5000);
    </script>
</body>
</html>
`
