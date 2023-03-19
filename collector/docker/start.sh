
/usr/bin/kindling-probe-loader

if [ -f "/opt/probe.o" ]; then
	export SYSDIG_BPF_PROBE="/opt/probe.o"
	/app/kindling-collector --config=/app/config/kindling-collector-config.yml
else
  echo "The kernel is not supported natively."
  echo "Please read FAQ http://kindling.harmonycloud.cn/docs/installation/faq/#error-precompiled-module-at-optkindling-is-not-found"
  sleep infinity
fi