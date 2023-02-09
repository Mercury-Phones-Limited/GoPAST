#!/usr/local/bin/bash

path=/usr/local/lib/asterisk/modules;
chgrp asterisk /usr/local/lib/asterisk;
chown root /usr/local/lib/asterisk;
chmod 550 /usr/local/lib/asterisk;
chgrp asterisk $path;
chown root $path;
chmod 550 $path;
chgrp wheel $path/*;
chown root $path/*;
chmod 440 $path/*;

###############
# App modules #
###############

chgrp asterisk $path/app_cdr.so;      # Writes ad hoc record to CDR.
chgrp asterisk $path/app_dial.so;     # Used to connect channels together (i.e., make phone calls).
chgrp asterisk $path/app_playback.so; # Plays pairs of tones of specified frequencies (DTMF mostly).
#chgrp asterisk $path/app_read.so;    # Requests input of digits from callers and assigns input to a variable.

##################
# Bridge modules #
##################

chgrp asterisk $path/bridge_simple.so; # Performs simple channel-to-channel bridging.

###############
# CDR modules #
###############

chgrp asterisk $path/cdr_adaptive_odbc.so; # Allows writing of CDRs through ODBC framework with ability to add custom fields.

###############
# CEL modules #
###############

###################
# Channel drivers #
###################

chgrp asterisk $path/chan_pjsip.so; # Session Initiation Protocol (SIP) channel driver.

#####################
# Codec Translators #
#####################

chgrp asterisk $path/codec_alaw.so; # A-law PCM codec used all over the world on the PSTN (except Canada/USA).
chgrp asterisk $path/codec_gsm.so;  # Global System for Mobile Communications (GSM) codec.

######################
# Dialplan Functions #
######################

chgrp asterisk $path/func_callerid.so; # Gets/sets caller ID.
chgrp asterisk $path/func_cdr.so;      # Gets/sets CDR variable.
chgrp asterisk $path/func_odbc.so;     # ODBC lookups.
#chgrp asterisk $path/func_rand.so;     # Random number dialplan function
#chgrp asterisk $path/func_sha1.so;     # SHA-1 computation dialplan function
chgrp asterisk $path/func_strings.so;  # Includes over a dozen string manipulation functions.
#chgrp asterisk $path/func_logic.so;   # Includes ISNULL(), SET(), EXISTS(), IF(), IFTIME(), and IMPORT() performs various logical functions.

#######################
# Format Interpreters #
#######################

chgrp asterisk $path/format_gsm.so;     # RPE-LTP (original GSM codec): .gsm
chgrp asterisk $path/format_wav.so;     # .wav
chgrp asterisk $path/format_wav_gsm.so; # GSM audio in a WAV container: .wav, .wav49

###############
# PBX Modules #
###############

chgrp asterisk $path/pbx_config.so;    # Provides the traditional dialplan language for Asterisk. Without this module, Asterisk cannot read extensions.conf.
#chgrp asterisk $path/pbx_realtime.so; # Realtime Switch - OBSOLETE! USE func_odbc.so

####################
# Resource Modules #
####################

########################
# res Calender modules #
########################

####################
# res odbc modules #
####################

chgrp asterisk $path/res_config_odbc.so;      # Pulls configuration information using ODBC.
chgrp asterisk $path/res_odbc.so;             # Provides common subroutines for other ODBC modules.
chgrp asterisk $path/res_odbc_transaction.so; # ODBC transaction resource (func_odbc.so dependancy).

#####################
# res other modules #
#####################

chgrp asterisk $path/res_clialiases.so; # This module provides the capability to create aliases to other CLI commands.

#####################
# res pjsip modules #
#####################

chgrp asterisk $path/res_pjproject.so; # PJPROJECT Log and Utility Support.

chgrp asterisk $path/res_pjsip.so;                               # Basic SIP resource.
chgrp asterisk $path/res_pjsip_authenticator_digest.so;          # PJSIP authentication resource. 
chgrp asterisk $path/res_pjsip_caller_id.so;                     # PJSIP Caller ID Support.
chgrp asterisk $path/res_pjsip_endpoint_identifier_anonymous.so; # PJSIP Anonymous endpoint identifier.
chgrp asterisk $path/res_pjsip_endpoint_identifier_ip.so;        # PJSIP IP endpoint identifier.
chgrp asterisk $path/res_pjsip_endpoint_identifier_user.so;      # PJSIP username endpoint identifier.
chgrp asterisk $path/res_pjsip_outbound_authenticator_digest.so; # PJSIP authentication resource.
chgrp asterisk $path/res_pjsip_outbound_publish.so;              # PJSIP Outbound Publish Support.
chgrp asterisk $path/res_pjsip_outbound_registration.so;         # PJSIP Outbound Registration Support.
chgrp asterisk $path/res_pjsip_pubsub.so;                        # PJSIP event resource.
chgrp asterisk $path/res_pjsip_registrar.so;                     # PJSIP Registrar Support.            
chgrp asterisk $path/res_pjsip_sdp_rtp.so;                       # PJSIP SDP RTP/AVP stream handler.
chgrp asterisk $path/res_pjsip_session.so;                       # PJSIP Session resource. 

chgrp asterisk $path/res_rtp_asterisk.so;     # Asterisk RTP Stack.
chgrp asterisk $path/res_srtp.so;             # Secure RTP (SRTP).

chgrp asterisk $path/res_sorcery_astdb.so;    # Sorcery Astdb Object Wizard. 
chgrp asterisk $path/res_sorcery_config.so;   # Sorcery Configuration File Object Wizard.
chgrp asterisk $path/res_sorcery_memory.so;   # Sorcery In-Memory Object Wizard.
chgrp asterisk $path/res_sorcery_realtime.so; # Sorcery Realtime Object Wizard.

chgrp asterisk $path/res_timing_timerfd.so; # Timerfd Timing Interface.
