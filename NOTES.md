**
Those notes were taken based on 
https://medium.com/@openmohan/dns-basics-and-building-simple-dns-server-in-go-6cb8e1cfe461
**

# Structure
- Basics
- Internal/External DNS Queries
- Zone Types and Resource Records
- Structure of queries and answers

# Basics of DNS:
- it helps the requester to find the IP address of a site
- Alternatively, we have host_files. But it'd be challenging to store many ip/sites into a single file

# Order of DNS Queries (lookup)
1. host_files (internal query)
2. DNS Cache (internal query)
3. DNS Server (external query)

# Authoritative Response
- Authoritative server: the server who has the info in its zone. Not in its cache but in a DNS zone file
- It is considered more reliable
- Example internal network: if I type something.corp.com, corp.com will be solved by a internal DNS server (same zone)

# Non-Authoritative Response
- When our DNS servers doesn't have an answer for a query, they'll ask the root hints server:
- Example:
CLIENT: careers.github.com?
DNS_SERVER: I don't know. Let me check with hint root server

DNS_SERVER: careers.github.com?
ROOT_SERVER(.): I don't know. But I can get you to .com DNS Server

DNS_SERVER: careers.github.com?
DOTCOM_SERVER(.com): I don't know. But I can get you to .github.com DNS Server

DNS_SERVER: careers.github.com?
GITHUB_SERVER(.github.com): Sure. IP is xxx.xxx.xxx.xxx, you can cache it for a while.

# Recursive vs Non-recursive query
- we can specify if a DNS server should pass the query to another DNS server
- we can use non-recursiveness to only get authoritative answers for example
- we can avoid DoS attacks

# Root Hints
- Data persisted in our OS
- Mapping of Root DNS Server name and their IP Address
- ICANN manages the root servers

# DNS Zone Types
- Forward lookup zone: Hostname -> IP Address
- Inverse lookup zone: IP Address -> Hostname
- Lookup zones contains resource records (RR) for serving DNS queries (Zone file)
- Zone files can be master file (authoritatively describes a zone) or may be used to list content of DNS cache

# Record Types
- A and AAAA: IPv4 and IPv6 host record
- CNAME (Canonical Name): provides alias name to any server. Example: subdomain docs.corp.com (CNAME) and corp.com (A)
- MX (Mail Exchange): Used for inter mail transfer. Example: send email from Gmail to Yahoo (Gmail asks a DNS for yahoo.com with type MX)
- PTR 
- NS

# Message Format
## Header
- ID (used to map questions with answers), 
- QR (query or response)
- OPCODE (type of query/op standard, inverse, status)
- AA (authoritative ans)
- TC (truncated message) 
- RD (recursion desired)
- RA (recursion available in response from DNS server)
- Z (reserved for future)
- RCODE (response code: 0-5)
- QDCOUNT (num of questions in req)
- ANCOUNT (num of answers in res)
- NSCOUNT (num of RR in authority records)
- ARCOUNT (num of RR in additional records section)
## Question
- QNAME (domain name we wish for resolving)
- QTYPE (A, AAAA, MX, TXT)
- QCLASS (IN, CH, HS)

## Answer, Authority and Additional (all three share the same format)
- NAME (domain name to which the RR pertains)
- TYPE (standard or inverse query)
- TTL (time-to-live in cache)
- RDLENGTH (response data length)
- RDATA (string of octets describing the resource. The format varies according to QTYPE and QCLASS)

# Building DNS server in Golang
- Generally, DNS queries and answers are lighweight UDP messages
- TCP is used when sending a large answer as DNS message (ex Zone Transfers, content is copied among DNS servers)
- By default DNS Server runs on port 53





