# Copyright 2011-2012 James McCauley
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at:
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

"""
An L2 learning switch.

It is derived from one written live for an SDN crash course.
It is somwhat similar to NOX's pyswitch in that it installs
exact-match rules for each flow.
"""

from pox.core import core
import pox.openflow.libopenflow_01 as of
from pox.lib.util import dpid_to_str, str_to_dpid
from pox.lib.util import str_to_bool
import time
import pox.lib.packet as pkt

log = core.getLogger()

# We don't want to flood immediately when a switch connects.
# Can be overriden on commandline.
_flood_delay = 0

class LearningSwitch (object):
  """
  The learning switch "brain" associated with a single OpenFlow switch.

  When we see a packet, we'd like to output it on a port which will
  eventually lead to the destination.  To accomplish this, we build a
  table that maps addresses to ports.

  We populate the table by observing traffic.  When we see a packet
  from some source coming from some port, we know that source is out
  that port.

  When we want to forward traffic, we look up the desintation in our
  table.  If we don't know the port, we simply send the message out
  all ports except the one it came in on.  (In the presence of loops,
  this is bad!).

  In short, our algorithm looks like this:

  For each packet from the switch:
  1) Use source address and switch port to update address/port table
  2) Is transparent = False and either Ethertype is LLDP or the packet's
     destination address is a Bridge Filtered address?
     Yes:
        2a) Drop packet -- don't forward link-local traffic (LLDP, 802.1x)
            DONE
  3) Is destination multicast?
     Yes:
        3a) Flood the packet
            DONE
  4) Port for destination address in our address/port table?
     No:
        4a) Flood the packet
            DONE
  5) Is output port the same as input port?
     Yes:
        5a) Drop packet and similar ones for a while
  6) Install flow table entry in the switch so that this
     flow goes out the appopriate port
     6a) Send the packet out appropriate port
  """
  def __init__ (self, connection, transparent):
    # Switch we'll be adding L2 learning switch capabilities to
    self.connection = connection
    self.transparent = transparent

    # Our table
    self.macToPort = {}

    # We want to hear PacketIn messages, so we listen
    # to the connection
    connection.addListeners(self)

    # We just use this to know when to log a helpful message
    self.hold_down_expired = _flood_delay == 0

    #log.debug("Initializing LearningSwitch, transparent=%s",
    #          str(self.transparent))

  def _handle_PacketIn (self, event):
    """
    Handle packet in messages from the switch to implement above algorithm.
    """

    packet = event.parsed
    #print(packet, "HOLA SOY UN PAKETE SIIIIII")
    def flood (message = None):
      """ Floods the packet """
      msg = of.ofp_packet_out()
      if time.time() - self.connection.connect_time >= _flood_delay:
        # Only flood if we've been connected for a little while...

        if self.hold_down_expired is False:
          # Oh yes it is!
          self.hold_down_expired = True
          log.info("%s: Flood hold-down expired -- flooding",
              dpid_to_str(event.dpid))

        if message is not None: log.debug(message)
        #log.debug("%i: flood %s -> %s", event.dpid,packet.src,packet.dst)
        # OFPP_FLOOD is optional; on some switches you may need to change
        # this to OFPP_ALL.
        msg.actions.append(of.ofp_action_output(port = of.OFPP_FLOOD))
      else:
        pass
        #log.info("Holding down flood for %s", dpid_to_str(event.dpid))
      msg.data = event.ofp
      msg.in_port = event.port
      self.connection.send(msg)

    def drop (duration = None):
      """
      Drops this packet and optionally installs a flow to continue
      dropping similar ones for a while
      """
      if duration is not None:
        if not isinstance(duration, tuple):
          duration = (duration,duration)
        msg = of.ofp_flow_mod()
        msg.match = of.ofp_match.from_packet(packet)
        msg.idle_timeout = duration[0]
        msg.hard_timeout = duration[1]
        msg.buffer_id = event.ofp.buffer_id
        self.connection.send(msg)
      elif event.ofp.buffer_id is not None:
        msg = of.ofp_packet_out()
        msg.buffer_id = event.ofp.buffer_id
        msg.in_port = event.port
        self.connection.send(msg)

    self.macToPort[packet.src] = event.port # 1
    """
    if not self.transparent: # 2
      
      if packet.type == packet.LLDP_TYPE or packet.dst.isBridgeFiltered():
        drop() # 2a
        return
    """
    if packet.dst.is_multicast:
      flood() # 3a
    else:
      """
      if packet.dst not in self.macToPort: # 4
        flood("Port for %s unknown -- flooding" % (packet.dst,)) # 4a
      else:
        port = self.macToPort[packet.dst]
      """
      #print(packet.dst)
      #port = self.macToPort[packet.dst]
      #print(port, "HOLA SOY EL PUERTO SIIIIIIII")
      """
      if port == event.port: # 5
        # 5a
        log.warning("Same port for packet from %s -> %s on %s.%s.  Drop."
            % (packet.src, packet.dst, dpid_to_str(event.dpid), port))
        drop(10)
        return
      """
      # 6
      #log.debug("installing flow for %s.%i -> %s.%i" %
      #          (packet.src, event.port, packet.dst, port))

      
      #print("AAAAAAAAAAAAA",packet.find("icmp"))
      papope = packet.find("tcp")
      print(packet)
      if papope != None: 
        if packet.find("arp") == None:
          if"HTTP" not in papope.payload and not(papope.SYN) and not(papope.ACK):
            print("Mensaje no es HTTP request")
            drop()
            return
      elif packet.find("arp") == None:
        print("Mensaje no es HTTP request")
        drop()
        return  

      
      msg = of.ofp_flow_mod()
      casita = ["00:00:00:00:00:01","00:00:00:00:00:02","00:00:00:00:00:03","00:00:00:00:00:04","00:00:00:00:00:05","00:00:00:00:00:06","00:00:00:00:00:07"]
      casita2 = ["00:00:00:00:00:01","00:00:00:00:00:02","00:00:00:00:00:03","00:00:00:00:00:04","00:00:00:00:00:05","00:00:00:00:00:06"]
      if str(packet.dst) not in casita or str(packet.src) not in casita:
        drop()
        return
      if str(packet.dst) in casita2 and str(packet.src) in casita2:
        drop()
        return
      msg.match = of.ofp_match.from_packet(packet, event.port)
      msg.idle_timeout = 10
      print ("Destino final: ",str(packet.dst),"Fuente: " ,str(packet.src))
      print ("Vengo desde:" ,event.port)
      msg.hard_timeout = 30

      
      if event.port == 29:
        msg.actions.append(of.ofp_action_output(port = 31))
        print("Voy a salir por:", 31)
      elif event.port == 28:
        msg.actions.append(of.ofp_action_output(port = 29))
        print("Voy a salir por:", 29)

      elif event.port == 32:
        if str(packet.dst) == "00:00:00:00:00:01":
          msg.actions.append(of.ofp_action_output(port = 1))
          print("Voy a salir por:", 29)

        elif str(packet.dst) == "00:00:00:00:00:02":
          msg.actions.append(of.ofp_action_output(port = 3))
          print("Voy a salir por:", 29)
        else:
          msg.actions.append(of.ofp_action_output(port = 5))
          print("Voy a salir por:", 5)
      elif event.port == 1 or event.port == 3:
          msg.actions.append(of.ofp_action_output(port = 7))
          print("Voy a salir por:", 7)

      elif event.port == 6:
        if str(packet.dst) == "00:00:00:00:00:03":
          msg.actions.append(of.ofp_action_output(port = 9))
          print("Voy a salir por:", 9)

        elif str(packet.dst) == "00:00:00:00:00:04":
          msg.actions.append(of.ofp_action_output(port = 11))
          print("Voy a salir por:", 11)
        else:
          msg.actions.append(of.ofp_action_output(port = 13))
          print("Voy a salir por:", 5)
      elif event.port == 9 or event.port == 11:
          msg.actions.append(of.ofp_action_output(port = 15))
          print("Voy a salir por:", 15)

      elif event.port == 14:
        if str(packet.dst) == "00:00:00:00:00:05":
          msg.actions.append(of.ofp_action_output(port = 17))
          print("Voy a salir por:", 17)

        elif str(packet.dst) == "00:00:00:00:00:06":
          msg.actions.append(of.ofp_action_output(port = 19))
          print("Voy a salir por:", 19)

      elif event.port == 17 or event.port == 19:
          msg.actions.append(of.ofp_action_output(port = 21))
          print("Voy a salir por:", 21)

      elif event.port in [22,24,26]:
        msg.actions.append(of.ofp_action_output(port = event.port+1))
        print("Voy a salir por:", event.port+1)

      elif event.port == 16:
        msg.actions.append(of.ofp_action_output(port = 25))
        print("Voy a salir por:", 25)
      elif event.port == 8:
        msg.actions.append(of.ofp_action_output(port = 27))
        print("Voy a salir por:", 27)

      else:
        msg.actions.append(of.ofp_action_output(port = port))
        print("Voy a salir por:", port) 
      msg.data = event.ofp # 6a
      
      self.connection.send(msg)

class l2_learning (object):
  """
  Waits for OpenFlow switches to connect and makes them learning switches.
  """
  def __init__ (self, transparent, ignore = None):
    """
    Initialize

    See LearningSwitch for meaning of 'transparent'
    'ignore' is an optional list/set of DPIDs to ignore
    """
    core.openflow.addListeners(self)
    self.transparent = transparent
    self.ignore = set(ignore) if ignore else ()

  def _handle_ConnectionUp (self, event):
    if event.dpid in self.ignore:
      log.debug("Ignoring connection %s" % (event.connection,))
      return
    log.debug("Connection %s" % (event.connection,))
    LearningSwitch(event.connection, self.transparent)


def launch (transparent=False, hold_down=_flood_delay, ignore = None):
  """
  Starts an L2 learning switch.
  """
  try:
    global _flood_delay
    _flood_delay = int(str(hold_down), 10)
    assert _flood_delay >= 0
  except:
    raise RuntimeError("Expected hold-down to be a number")

  if ignore:
    ignore = ignore.replace(',', ' ').split()
    ignore = set(str_to_dpid(dpid) for dpid in ignore)

  core.registerNew(l2_learning, str_to_bool(transparent), ignore)

